package sqlite3

/*
// make go-sqlite3 use embedded library without code changes
#cgo CFLAGS: -DUSE_LIBSQLITE3

// enable encryption codec in sqlite
#cgo CFLAGS: -DSQLITE_HAS_CODEC

// use memory for temporay storage in sqlite
#cgo CFLAGS: -DSQLITE_TEMP_STORE=2

// use platform-native crypto: CommonCrypto on macOS, OpenSSL on Linux/Windows
#cgo darwin CFLAGS: -DSQLCIPHER_CRYPTO_CC
#cgo darwin LDFLAGS: -framework Security
#cgo linux CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL
#cgo linux LDFLAGS: -lcrypto
#cgo windows CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL
#cgo windows LDFLAGS: -lcrypto

// sqlcipher extra init/shutdown hooks (required by SQLCipher 4.6+)
#cgo CFLAGS: -DSQLITE_EXTRA_INIT=sqlcipher_extra_init
#cgo CFLAGS: -DSQLITE_EXTRA_SHUTDOWN=sqlcipher_extra_shutdown

// enable FTS5 full-text search
#cgo CFLAGS: -DSQLITE_ENABLE_FTS5

// FTS5 needs libm. Upstream enables FTS5 only under the sqlite_fts5 build tag,
// and sqlite3_opt_fts5.go pairs the CFLAG above with `#cgo LDFLAGS: -lm`. This
// fork turns FTS5 on unconditionally here and did not bring the link flag
// across, so fts5Bm25GetData's call to log() had nothing to resolve against:
//
//	sqlite3.c: undefined reference to `log'
//
// Linux only. macOS has the maths functions in libSystem, and the Windows build
// already links -lmingwex in sqlite3_windows.go, which provides them — so those
// two hid the omission, and only a plain `go build` on Linux ever showed it.
#cgo linux LDFLAGS: -lm

// disable assertions
#cgo CFLAGS: -DNDEBUG

// set operating specific sqlite flags
#cgo linux CFLAGS: -DSQLITE_OS_UNIX=1
#cgo windows CFLAGS: -DSQLITE_OS_WIN=1
*/
import "C"
