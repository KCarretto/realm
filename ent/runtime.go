// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/kcarretto/realm/ent/credential"
	"github.com/kcarretto/realm/ent/file"
	"github.com/kcarretto/realm/ent/implant"
	"github.com/kcarretto/realm/ent/implantcallbackconfig"
	"github.com/kcarretto/realm/ent/implantconfig"
	"github.com/kcarretto/realm/ent/implantserviceconfig"
	"github.com/kcarretto/realm/ent/schema"
	"github.com/kcarretto/realm/ent/target"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	credentialFields := schema.Credential{}.Fields()
	_ = credentialFields
	// credentialDescPrincipal is the schema descriptor for principal field.
	credentialDescPrincipal := credentialFields[0].Descriptor()
	// credential.PrincipalValidator is a validator for the "principal" field. It is called by the builders before save.
	credential.PrincipalValidator = credentialDescPrincipal.Validators[0].(func(string) error)
	// credentialDescSecret is the schema descriptor for secret field.
	credentialDescSecret := credentialFields[1].Descriptor()
	// credential.SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	credential.SecretValidator = func() func(string) error {
		validators := credentialDescSecret.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(secret string) error {
			for _, fn := range fns {
				if err := fn(secret); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescName is the schema descriptor for name field.
	fileDescName := fileFields[0].Descriptor()
	// file.NameValidator is a validator for the "name" field. It is called by the builders before save.
	file.NameValidator = fileDescName.Validators[0].(func(string) error)
	// fileDescSize is the schema descriptor for size field.
	fileDescSize := fileFields[1].Descriptor()
	// file.DefaultSize holds the default value on creation for the size field.
	file.DefaultSize = fileDescSize.Default.(int)
	// file.SizeValidator is a validator for the "size" field. It is called by the builders before save.
	file.SizeValidator = fileDescSize.Validators[0].(func(int) error)
	// fileDescHash is the schema descriptor for hash field.
	fileDescHash := fileFields[2].Descriptor()
	// file.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	file.HashValidator = fileDescHash.Validators[0].(func(string) error)
	// fileDescCreatedAt is the schema descriptor for createdAt field.
	fileDescCreatedAt := fileFields[3].Descriptor()
	// file.DefaultCreatedAt holds the default value on creation for the createdAt field.
	file.DefaultCreatedAt = fileDescCreatedAt.Default.(func() time.Time)
	implantFields := schema.Implant{}.Fields()
	_ = implantFields
	// implantDescSessionID is the schema descriptor for sessionID field.
	implantDescSessionID := implantFields[0].Descriptor()
	// implant.SessionIDValidator is a validator for the "sessionID" field. It is called by the builders before save.
	implant.SessionIDValidator = implantDescSessionID.Validators[0].(func(string) error)
	implantcallbackconfigFields := schema.ImplantCallbackConfig{}.Fields()
	_ = implantcallbackconfigFields
	// implantcallbackconfigDescURI is the schema descriptor for uri field.
	implantcallbackconfigDescURI := implantcallbackconfigFields[0].Descriptor()
	// implantcallbackconfig.URIValidator is a validator for the "uri" field. It is called by the builders before save.
	implantcallbackconfig.URIValidator = implantcallbackconfigDescURI.Validators[0].(func(string) error)
	// implantcallbackconfigDescPriority is the schema descriptor for priority field.
	implantcallbackconfigDescPriority := implantcallbackconfigFields[2].Descriptor()
	// implantcallbackconfig.DefaultPriority holds the default value on creation for the priority field.
	implantcallbackconfig.DefaultPriority = implantcallbackconfigDescPriority.Default.(int)
	// implantcallbackconfig.PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	implantcallbackconfig.PriorityValidator = implantcallbackconfigDescPriority.Validators[0].(func(int) error)
	// implantcallbackconfigDescTimeout is the schema descriptor for timeout field.
	implantcallbackconfigDescTimeout := implantcallbackconfigFields[3].Descriptor()
	// implantcallbackconfig.DefaultTimeout holds the default value on creation for the timeout field.
	implantcallbackconfig.DefaultTimeout = implantcallbackconfigDescTimeout.Default.(int)
	// implantcallbackconfig.TimeoutValidator is a validator for the "timeout" field. It is called by the builders before save.
	implantcallbackconfig.TimeoutValidator = implantcallbackconfigDescTimeout.Validators[0].(func(int) error)
	// implantcallbackconfigDescInterval is the schema descriptor for interval field.
	implantcallbackconfigDescInterval := implantcallbackconfigFields[4].Descriptor()
	// implantcallbackconfig.DefaultInterval holds the default value on creation for the interval field.
	implantcallbackconfig.DefaultInterval = implantcallbackconfigDescInterval.Default.(int)
	// implantcallbackconfig.IntervalValidator is a validator for the "interval" field. It is called by the builders before save.
	implantcallbackconfig.IntervalValidator = implantcallbackconfigDescInterval.Validators[0].(func(int) error)
	// implantcallbackconfigDescJitter is the schema descriptor for jitter field.
	implantcallbackconfigDescJitter := implantcallbackconfigFields[5].Descriptor()
	// implantcallbackconfig.DefaultJitter holds the default value on creation for the jitter field.
	implantcallbackconfig.DefaultJitter = implantcallbackconfigDescJitter.Default.(int)
	// implantcallbackconfig.JitterValidator is a validator for the "jitter" field. It is called by the builders before save.
	implantcallbackconfig.JitterValidator = implantcallbackconfigDescJitter.Validators[0].(func(int) error)
	implantconfigFields := schema.ImplantConfig{}.Fields()
	_ = implantconfigFields
	// implantconfigDescName is the schema descriptor for name field.
	implantconfigDescName := implantconfigFields[0].Descriptor()
	// implantconfig.NameValidator is a validator for the "name" field. It is called by the builders before save.
	implantconfig.NameValidator = implantconfigDescName.Validators[0].(func(string) error)
	// implantconfigDescAuthToken is the schema descriptor for authToken field.
	implantconfigDescAuthToken := implantconfigFields[1].Descriptor()
	// implantconfig.DefaultAuthToken holds the default value on creation for the authToken field.
	implantconfig.DefaultAuthToken = implantconfigDescAuthToken.Default.(func() string)
	implantserviceconfigFields := schema.ImplantServiceConfig{}.Fields()
	_ = implantserviceconfigFields
	// implantserviceconfigDescName is the schema descriptor for name field.
	implantserviceconfigDescName := implantserviceconfigFields[0].Descriptor()
	// implantserviceconfig.NameValidator is a validator for the "name" field. It is called by the builders before save.
	implantserviceconfig.NameValidator = implantserviceconfigDescName.Validators[0].(func(string) error)
	// implantserviceconfigDescDescription is the schema descriptor for description field.
	implantserviceconfigDescDescription := implantserviceconfigFields[1].Descriptor()
	// implantserviceconfig.DefaultDescription holds the default value on creation for the description field.
	implantserviceconfig.DefaultDescription = implantserviceconfigDescDescription.Default.(string)
	// implantserviceconfigDescExecutablePath is the schema descriptor for executablePath field.
	implantserviceconfigDescExecutablePath := implantserviceconfigFields[2].Descriptor()
	// implantserviceconfig.ExecutablePathValidator is a validator for the "executablePath" field. It is called by the builders before save.
	implantserviceconfig.ExecutablePathValidator = implantserviceconfigDescExecutablePath.Validators[0].(func(string) error)
	targetFields := schema.Target{}.Fields()
	_ = targetFields
	// targetDescName is the schema descriptor for name field.
	targetDescName := targetFields[0].Descriptor()
	// target.NameValidator is a validator for the "name" field. It is called by the builders before save.
	target.NameValidator = func() func(string) error {
		validators := targetDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// targetDescForwardConnectIP is the schema descriptor for forwardConnectIP field.
	targetDescForwardConnectIP := targetFields[1].Descriptor()
	// target.ForwardConnectIPValidator is a validator for the "forwardConnectIP" field. It is called by the builders before save.
	target.ForwardConnectIPValidator = targetDescForwardConnectIP.Validators[0].(func(string) error)
}
