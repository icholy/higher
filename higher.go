package higher

import (
	"reflect"
	"sync"
)

func sliceFilter(inValue, fnValue reflect.Value) reflect.Value {
	var (
		inType     = inValue.Type()
		inValueLen = inValue.Len()
		outValue   = reflect.MakeSlice(inType, 0, 1)
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		if fnValue.Call(args)[0].Bool() {
			outValue = reflect.Append(outValue, args[0])
		}
	}
	return outValue
}

func Filter(in, fn interface{}) interface{} {
	return sliceFilter(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	).Interface()
}

func sliceMap(inValue, fnValue reflect.Value) reflect.Value {
	var (
		inValueLen = inValue.Len()
		fnOutType  = fnValue.Type().Out(0)
		outType    = reflect.SliceOf(fnOutType)
		outValue   = reflect.MakeSlice(outType, inValueLen, inValueLen)
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		outValue.Index(i).Set(fnValue.Call(args)[0])
	}
	return outValue
}

func sliceParallelMap(inValue, fnValue reflect.Value) reflect.Value {
	var (
		inValueLen = inValue.Len()
		fnOutType  = fnValue.Type().Out(0)
		outType    = reflect.SliceOf(fnOutType)
		outValue   = reflect.MakeSlice(outType, inValueLen, inValueLen)
		wg         sync.WaitGroup
	)
	wg.Add(inValueLen)
	for i := 0; i < inValueLen; i++ {
		go func(j int) {
			args := []reflect.Value{inValue.Index(j)}
			outValue.Index(j).Set(fnValue.Call(args)[0])
			wg.Done()
		}(i)
	}
	wg.Wait()
	return outValue
}

func Map(in, fn interface{}) interface{} {
	return sliceMap(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	).Interface()
}

func PMap(in, fn interface{}) interface{} {
	return sliceParallelMap(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	).Interface()
}

func sliceReduce(inValue, fnValue, accValue reflect.Value) reflect.Value {
	var (
		inValueLen = inValue.Len()
		args       = make([]reflect.Value, 2)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = accValue
		args[1] = inValue.Index(i)
		accValue = fnValue.Call(args)[0]
	}
	return accValue
}

func Reduce(in, fn, acc interface{}) interface{} {
	return sliceReduce(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
		reflect.ValueOf(acc),
	).Interface()
}

func sliceForEach(inValue, fnValue reflect.Value) {
	var (
		inValueLen = inValue.Len()
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		_ = fnValue.Call(args)
	}
}

func ForEach(in, fn interface{}) {
	sliceForEach(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	)
}

func sliceTap(inValue, fnValue reflect.Value) reflect.Value {
	var (
		inValueLen = inValue.Len()
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		_ = fnValue.Call(args)
	}
	return inValue
}

func Tap(in, fn interface{}) interface{} {
	return sliceTap(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	).Interface()
}

func sliceAny(inValue, fnValue reflect.Value) bool {
	var (
		inValueLen = inValue.Len()
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		if fnValue.Call(args)[0].Bool() {
			return true
		}
	}
	return false
}

func Any(in, fn interface{}) bool {
	return sliceAny(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	)
}

func sliceEvery(inValue, fnValue reflect.Value) bool {
	var (
		inValueLen = inValue.Len()
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		if !fnValue.Call(args)[0].Bool() {
			return false
		}
	}
	return true
}

func Every(in, fn interface{}) bool {
	return sliceEvery(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	)
}

func sliceContains(inValue reflect.Value, v interface{}) bool {
	var (
		inValueLen = inValue.Len()
	)
	for i := 0; i < inValueLen; i++ {
		if reflect.DeepEqual(v, inValue.Index(i).Interface()) {
			return true
		}
	}
	return false
}

func Contains(in, v interface{}) bool {
	return sliceContains(
		reflect.ValueOf(in), v,
	)
}

func sliceFind(inValue, fnValue reflect.Value) reflect.Value {
	var (
		inValueLen = inValue.Len()
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		if fnValue.Call(args)[0].Bool() {
			return args[0]
		}
	}
	return reflect.Zero(inValue.Type().Elem())
}

func Find(in, fn interface{}) interface{} {
	return sliceFind(
		reflect.ValueOf(in),
		reflect.ValueOf(fn),
	).Interface()
}

type Wrapped struct {
	value reflect.Value
}

func Wrap(in interface{}) Wrapped {
	return Wrapped{reflect.ValueOf(in)}
}

func (w Wrapped) Map(fn interface{}) Wrapped {
	return Wrapped{
		value: sliceMap(w.value, reflect.ValueOf(fn)),
	}
}

func (w Wrapped) PMap(fn interface{}) Wrapped {
	return Wrapped{
		value: sliceParallelMap(w.value, reflect.ValueOf(fn)),
	}
}

func (w Wrapped) Filter(fn interface{}) Wrapped {
	return Wrapped{
		value: sliceFilter(w.value, reflect.ValueOf(fn)),
	}
}

func (w Wrapped) Reduce(fn interface{}, acc interface{}) Wrapped {
	return Wrapped{
		value: sliceReduce(
			w.value,
			reflect.ValueOf(fn),
			reflect.ValueOf(acc),
		),
	}
}

func (w Wrapped) ForEach(fn interface{}) {
	sliceForEach(w.value, reflect.ValueOf(fn))
}

func (w Wrapped) Tap(fn interface{}) Wrapped {
	return Wrapped{
		value: sliceTap(
			w.value,
			reflect.ValueOf(fn),
		),
	}
}

func (w Wrapped) Any(fn interface{}) bool {
	return sliceAny(w.value, reflect.ValueOf(fn))
}

func (w Wrapped) Every(fn interface{}) bool {
	return sliceEvery(w.value, reflect.ValueOf(fn))
}

func (w Wrapped) Contains(v interface{}) bool {
	return sliceContains(w.value, v)
}

func (w Wrapped) Find(fn interface{}) interface{} {
	return sliceFind(w.value, reflect.ValueOf(fn)).Interface()
}

func (w Wrapped) Val() interface{} {
	return w.value.Interface()
}
