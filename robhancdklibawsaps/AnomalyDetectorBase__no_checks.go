//go:build no_runtime_type_checking

package robhancdklibawsaps

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnomalyDetectorBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AnomalyDetectorBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AnomalyDetectorBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAnomalyDetectorBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAnomalyDetectorBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAnomalyDetectorBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAnomalyDetectorBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

