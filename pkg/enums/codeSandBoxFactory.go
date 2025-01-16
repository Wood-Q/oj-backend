package enums

type CodeSandboxFactory struct{}

func (f *CodeSandboxFactory) NewInstance(sandboxType string) CodeSandbox {
	switch sandboxType {
	case "example":
		return &ExampleCodeSandbox{}
	case "remote":
		return &RemoteCodeSandbox{}
	case "thirdParty":
		return &ThirdPartyCodeSandbox{}
	default:
		return nil
	}
}
