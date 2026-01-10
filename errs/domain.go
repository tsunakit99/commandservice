package errs

// ドメインに関するエラー型
type DomainError struct {
	message string
}

// エラーメッセージを返す
func (e *DomainError) Error() string {
	return e.message
}

// ドメインエラーを生成するコンストラクタ
func NewDomainError(message string) *DomainError {
	return &DomainError{message: message}
}
