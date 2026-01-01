//go:build no_runtime_type_checking

package robhancdklibawsaps

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourcePolicyBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicyBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicyBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateResourcePolicyBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResourcePolicyBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateResourcePolicyBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewResourcePolicyBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

