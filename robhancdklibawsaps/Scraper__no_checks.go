//go:build no_runtime_type_checking

package robhancdklibawsaps

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Scraper) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Scraper) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Scraper) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_Scraper) validateGetScraperIdParameters(scraperArn *string) error {
	return nil
}

func validateScraper_FrommScraperAttributesParameters(scope constructs.Construct, id *string, attrs *ScraperAttributes) error {
	return nil
}

func validateScraper_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScraper_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateScraper_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateScraper_IsScraperParameters(x interface{}) error {
	return nil
}

func validateNewScraperParameters(scope constructs.Construct, id *string, props *ScraperProps) error {
	return nil
}

