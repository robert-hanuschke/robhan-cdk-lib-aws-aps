//go:build no_runtime_type_checking

package robhancdklibawsaps

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScraperBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ScraperBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ScraperBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_ScraperBase) validateGetScraperIdParameters(scraperArn *string) error {
	return nil
}

func validateScraperBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScraperBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateScraperBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewScraperBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

