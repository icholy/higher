package higher

import "reflect"

func Filter(in interface{}, fn interface{}) interface{} {
	var (
		inType     = reflect.TypeOf(in)
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
		outValue   = reflect.MakeSlice(inType, 0, 1)
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		if fnValue.Call(args)[0].Bool() {
			outValue = reflect.Append(outValue, args[0])
		}
	}
	return outValue.Interface()
}

func Map(in interface{}, fn interface{}) interface{} {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
		fnOutType  = reflect.TypeOf(fn).Out(0)
		outType    = reflect.SliceOf(fnOutType)
		outValue   = reflect.MakeSlice(outType, inValueLen, inValueLen)
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		rets := fnValue.Call(args)
		outValue.Index(i).Set(rets[0])
	}
	return outValue.Interface()
}

func Reduce(in interface{}, fn interface{}, acc interface{}) interface{} {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
		accValue   = reflect.ValueOf(acc)
		args       = make([]reflect.Value, 2)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = accValue
		args[1] = inValue.Index(i)
		accValue = fnValue.Call(args)[0]
	}
	return accValue.Interface()
}

func ForEach(in interface{}, fn interface{}) {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		_ = fnValue.Call(args)
	}
}

func Tap(in interface{}, fn interface{}) interface{} {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		_ = fnValue.Call(args)
	}
	return in
}

func Any(in interface{}, fn interface{}) bool {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
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

func Every(in interface{}, fn interface{}) bool {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
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

func Contains(in interface{}, v interface{}) bool {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
	)
	for i := 0; i < inValueLen; i++ {
		if reflect.DeepEqual(v, inValue.Index(i).Interface()) {
			return true
		}
	}
	return false
}

func Find(in interface{}, fn interface{}) interface{} {
	var (
		inValue    = reflect.ValueOf(in)
		inValueLen = inValue.Len()
		fnValue    = reflect.ValueOf(fn)
		args       = make([]reflect.Value, 1)
	)
	for i := 0; i < inValueLen; i++ {
		args[0] = inValue.Index(i)
		if fnValue.Call(args)[0].Bool() {
			return args[0].Interface()
		}
	}
	return nil
}

type Wrapped struct {
	value interface{}
}

func Wrap(in interface{}) Wrapped {
	return Wrapped{in}
}

func (w Wrapped) Map(fn interface{}) Wrapped {
	return Wrapped{Map(w.value, fn)}
}

func (w Wrapped) Filter(fn interface{}) Wrapped {
	return Wrapped{Filter(w.value, fn)}
}

func (w Wrapped) Reduce(fn interface{}, acc interface{}) Wrapped {
	return Wrapped{Reduce(w.value, fn, acc)}
}

func (w Wrapped) ForEach(fn interface{}) {
	ForEach(w.value, fn)
}

func (w Wrapped) Tap(fn interface{}) Wrapped {
	return Wrapped{Tap(w.value, fn)}
}

func (w Wrapped) Any(fn interface{}) bool {
	return Any(w.value, fn)
}

func (w Wrapped) Every(fn interface{}) bool {
	return Every(w.value, fn)
}

func (w Wrapped) Contains(v interface{}) bool {
	return Contains(w.value, v)
}

func (w Wrapped) Find(fn interface{}) interface{} {
	return Find(w.value, fn)
}

func (w Wrapped) Val() interface{} {
	return w.value
}
