//go:build no_runtime_type_checking

package robhancdklibawsaps

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RuleGroupsNamespaceBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RuleGroupsNamespaceBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RuleGroupsNamespaceBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRuleGroupsNamespaceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRuleGroupsNamespaceBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRuleGroupsNamespaceBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRuleGroupsNamespaceBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

