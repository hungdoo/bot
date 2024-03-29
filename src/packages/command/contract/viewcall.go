package contract

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

type ViewCall struct {
	method    string
	arguments []interface{}
}

func NewViewCall(method string, arguments []interface{}) ViewCall {
	return ViewCall{
		method:    method,
		arguments: arguments,
	}

}

func (call ViewCall) Validate() error {
	if _, err := call.argsCallData(); err != nil {
		return err
	}
	return nil
}

var insideParens = regexp.MustCompile("\\(.*?\\)")
var numericArg = regexp.MustCompile("u?int(256)|(8)")

func (call ViewCall) ArgumentTypes() []string {
	rawArgs := insideParens.FindAllString(call.method, -1)[0]
	rawArgs = strings.Replace(rawArgs, "(", "", -1)
	rawArgs = strings.Replace(rawArgs, ")", "", -1)
	if rawArgs == "" {
		return []string{}
	}
	args := strings.Split(rawArgs, ",")
	for index, arg := range args {
		args[index] = strings.Trim(arg, " ")
	}
	return args
}

func (call ViewCall) ReturnTypes() []string {
	rawArgs := insideParens.FindAllString(call.method, -1)[1]
	rawArgs = strings.Replace(rawArgs, "(", "", -1)
	rawArgs = strings.Replace(rawArgs, ")", "", -1)
	args := strings.Split(rawArgs, ",")
	for index, arg := range args {
		args[index] = strings.Trim(arg, " ")
	}
	return args
}

func (call ViewCall) CallData() ([]byte, error) {
	argsSuffix, err := call.argsCallData()
	if err != nil {
		return nil, err
	}
	methodPrefix, err := call.methodCallData()
	if err != nil {
		return nil, err
	}

	payload := make([]byte, 0)
	payload = append(payload, methodPrefix...)
	payload = append(payload, argsSuffix...)

	return payload, nil
}

func (call ViewCall) methodCallData() ([]byte, error) {
	methodParts := strings.Split(call.method, ")(")
	var method string
	if len(methodParts) > 1 {
		method = fmt.Sprintf("%s)", methodParts[0])
	} else {
		method = methodParts[0]
	}
	hash := crypto.Keccak256([]byte(method))
	return hash[0:4], nil
}

func (call ViewCall) argsCallData() ([]byte, error) {
	argTypes := call.ArgumentTypes()
	if len(argTypes) != len(call.arguments) {
		return nil, fmt.Errorf("number of argument types doesn't match with number of arguments for method %s", call.method)
	}
	argumentValues := make([]interface{}, len(call.arguments))
	arguments := make(abi.Arguments, len(call.arguments))

	for index, argTypeStr := range argTypes {
		argType, err := abi.NewType(argTypeStr, "", nil)
		if err != nil {
			return nil, err
		}

		arguments[index] = abi.Argument{Type: argType}
		argumentValues[index], err = call.getArgument(index, argTypeStr)
		if err != nil {
			return nil, err
		}
	}

	return arguments.Pack(argumentValues...)
}

func (call ViewCall) getArgument(index int, argumentType string) (interface{}, error) {
	arg := call.arguments[index]
	if argumentType == "address" {
		address, ok := arg.(string)
		if !ok {
			return nil, fmt.Errorf("expected address argument to be a string")
		}
		return toByteArray(address)
	} else if argumentType == "bool" {
		argStr, ok := arg.(string)
		if !ok {
			return nil, fmt.Errorf("expected bool argument to be a string: true/false")
		}

		return strconv.ParseBool(argStr)
	} else if numericArg.MatchString(argumentType) {
		if num, ok := arg.(json.Number); ok {
			if v, err := num.Int64(); err != nil {
				return big.NewInt(v), nil
			} else if v, err := num.Float64(); err != nil {
				return big.NewInt(int64(v)), nil
			} else {
			}
		} else {
			int64 := reflect.TypeOf(int64(0))
			argType := reflect.TypeOf(arg)
			kind := argType.Kind()
			if kind == reflect.String {
				if val, ok := new(big.Int).SetString(call.arguments[index].(string), 10); !ok {
					return nil, fmt.Errorf("could not parse %s as a base 10 number", call.arguments[index])
				} else {
					return val, nil
				}
			} else if argType.ConvertibleTo(int64) {
				return big.NewInt(reflect.ValueOf(arg).Convert(int64).Int()), nil
			}
		}
	}
	return arg, nil
}

func (call ViewCall) Decode(raw []byte) ([]interface{}, error) {
	retTypes := call.ReturnTypes()
	args := make(abi.Arguments, 0)
	for index, retTypeStr := range retTypes {
		retType, err := abi.NewType(retTypeStr, "", nil)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{Name: fmt.Sprintf("ret%d", index), Type: retType})
	}
	decoded := make(map[string]interface{})
	err := args.UnpackIntoMap(decoded, raw)
	if err != nil {
		return nil, err
	}
	returns := make([]interface{}, len(retTypes))
	for index := range retTypes {
		key := fmt.Sprintf("ret%d", index)
		returns[index] = decoded[key]
	}
	return returns, nil
}

func toByteArray(address string) ([20]byte, error) {
	var addressBytes [20]byte
	address = strings.Replace(address, "0x", "", -1)
	addressBytesSlice, err := hex.DecodeString(address)
	if err != nil {
		return addressBytes, err
	}

	copy(addressBytes[:], addressBytesSlice[:])
	return addressBytes, nil
}
