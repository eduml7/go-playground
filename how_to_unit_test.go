package main

//To run tests in go: $go test
import (
	"fmt"
	"testing"
	"time"

	//Testify is a lib who gives us a lot of utilities for testing,
	//expanding the ones in testing native lib
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//You can define a single test and run it
func TestWithoutSuite_1(t *testing.T) {
	//This function allows go to execute the tests in parallel, instead of sequentially.
	//Only tests with this function applied are executed in parallel
	t.Parallel()
	time.Sleep(1 * time.Second)
	expected := "value"
	result := "value"
	fmt.Println("Launching TestWithoutSuite_1")
	assert.Equal(t, expected, result, "")
}
func TestWithoutSuite_2(t *testing.T) {
	t.Parallel()
	expected := "value"
	result := "value"
	fmt.Println("Launching TestWithoutSuite_2")
	assert.Equal(t, expected, result, "")
}

//But if you want to make a test structure consistent you can define a test suite,
//because all _test files are considered test suites. Whith this structure you can
//package a bunch of tests in a single suite.
//
//In addition, you have this  pipeline methods for the suite:
// BeforeTest(suiteName, testName string) - Runs right before the test starts
// AfterTest(suiteName, testName string) - Runs right after the test finishes
// SetupSuite() - Runs before the tests in the suite
// SetupTest() - Runs before each test in the suite
// TearDownTest() - Runs after each test in the suite
// TearDownSuite() - Runs after all the tests in the suite have been run

type FinderTestSuite struct {
	suite.Suite
}

func (suite *FinderTestSuite) TestWithSuiteAndAssertT() {
	expected := "value"
	result := "value"
	fmt.Println("Launching TestWithSuiteAndAssertT")
	//Here access to the *testing.T variable inside the suite, with suite.T()
	assert.Equal(suite.T(), expected, result, "")
}

func (suite *FinderTestSuite) TestWithSuite() {
	expected := "value"
	result := "value"
	fmt.Println("Launching TestWithSuite")
	suite.Equal(expected, result)
}

//Runs automatically by $ go test
//You must run the test suite inside of it
func TestSuite(t *testing.T) {
	suite.Run(t, new(FinderTestSuite))
}
