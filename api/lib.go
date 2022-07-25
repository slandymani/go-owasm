package api

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR} -lgo_owasm
// #include "odin_bindings.h"
//
// typedef Span (*get_calldata_fn)(env_t*);
// Span cGetCalldata_cgo(env_t *e);
// typedef void (*set_return_data_fn)(env_t*, Span);
// void cSetReturnData_cgo(env_t *e, Span data);
// typedef int64_t (*get_ask_count_fn)(env_t*);
// int64_t cGetAskCount_cgo(env_t *e);
// typedef int64_t (*get_min_count_fn)(env_t*);
// int64_t cGetMinCount_cgo(env_t *e);
// typedef int64_t (*get_prepare_time_fn)(env_t*);
// int64_t cGetPrepareTime_cgo(env_t *e);
// typedef int64_t (*get_execute_time_fn)(env_t*);
// int64_t cGetExecuteTime_cgo(env_t *e);
// typedef int64_t (*get_ans_count_fn)(env_t*);
// int64_t cGetAnsCount_cgo(env_t *e);
// typedef void (*ask_external_data_fn)(env_t*, int64_t eid, int64_t did);
// void cAskExternalData_cgo(env_t *e, int64_t eid, int64_t did);
// typedef int64_t (*get_external_data_status_fn)(env_t*, int64_t eid, int64_t vid);
// int64_t cGetExternalDataStatus_cgo(env_t *e, int64_t eid, int64_t vid);
// typedef Span (*get_external_data_fn)(env_t*, int64_t eid, int64_t vid);
// Span cGetExternalData_cgo(env_t *e, int64_t eid, int64_t vid);
import "C"
import (
	"unsafe"
)

type RunOutput struct {
	GasUsed uint32
}

type Vm struct {
	cache Cache
}

func NewVm(size uint32) (*Vm, error) {
	cache, err := InitCache(size)

	if err != nil {
		return nil, err
	}

	return &Vm{
		cache: cache,
	}, nil
}

func (vm Vm) Compile(code []byte, spanSize int) ([]byte, error) {
	inputSpan := copySpan(code)
	defer freeSpan(inputSpan)
	outputSpan := newSpan(spanSize)
	defer freeSpan(outputSpan)
	err := toGoError(C.do_compile(inputSpan, &outputSpan))
	return readSpan(outputSpan), err
}

func (vm Vm) Prepare(code []byte, gasLimit uint32, spanSize int64, env EnvInterface) (RunOutput, error) {
	return vm.run(code, gasLimit, spanSize, true, env)
}

func (vm Vm) Execute(code []byte, gasLimit uint32, spanSize int64, env EnvInterface) (RunOutput, error) {
	return vm.run(code, gasLimit, spanSize, false, env)
}

func (vm Vm) run(code []byte, gasLimit uint32, spanSize int64, isPrepare bool, env EnvInterface) (RunOutput, error) {
	codeSpan := copySpan(code)
	defer freeSpan(codeSpan)
	envIntl := createEnvIntl(env)
	output := C.RunOutput{}
	err := toGoError(C.do_run(vm.cache.ptr, codeSpan, C.uint32_t(gasLimit), C.int64_t(spanSize), C.bool(isPrepare), C.Env{
		env: (*C.env_t)(unsafe.Pointer(envIntl)),
		dis: C.EnvDispatcher{
			get_calldata:             C.get_calldata_fn(C.cGetCalldata_cgo),
			set_return_data:          C.set_return_data_fn(C.cSetReturnData_cgo),
			get_ask_count:            C.get_ask_count_fn(C.cGetAskCount_cgo),
			get_min_count:            C.get_min_count_fn(C.cGetMinCount_cgo),
			get_prepare_time:         C.get_prepare_time_fn(C.cGetPrepareTime_cgo),
			get_execute_time:         C.get_execute_time_fn(C.cGetExecuteTime_cgo),
			get_ans_count:            C.get_ans_count_fn(C.cGetAnsCount_cgo),
			ask_external_data:        C.ask_external_data_fn(C.cAskExternalData_cgo),
			get_external_data_status: C.get_external_data_status_fn(C.cGetExternalDataStatus_cgo),
			get_external_data:        C.get_external_data_fn(C.cGetExternalData_cgo),
		},
	}, &output))
	if err != nil {
		return RunOutput{}, err
	} else {
		return RunOutput{GasUsed: uint32(output.gas_used)}, nil
	}
}

type Cache struct {
	ptr *C.cache_t
}

func InitCache(cacheSize uint32) (Cache, error) {
	ptr, err := C.oracle_init_cache(C.uint32_t(cacheSize))
	if err != nil {
		return Cache{}, err
	}
	return Cache{ptr: ptr}, nil
}

func ReleaseCache(cache Cache) {
	C.oracle_release_cache(cache.ptr)
}
