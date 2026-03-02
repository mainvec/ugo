package assert


func Assert(assertion bool, assertionFailedMessage string) {
	if !assertion {
		panic (assertionFailedMessage)
	}
}
