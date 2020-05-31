use crate::span::Span;

#[repr(C)]
pub struct env_t {
    _private: [u8; 0],
}

#[repr(C)]
pub struct EnvDispatcher {
    pub get_ask_count: extern "C" fn(*mut env_t) -> i64,
}

#[repr(C)]
pub struct Env {
    pub env: *mut env_t,
    pub dis: EnvDispatcher,
}
