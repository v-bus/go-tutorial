//Package write2file describes write file examples
package write2file

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"os"
	"reflect"
)

const ftdIncomAttrErr = "Incoming arguments are invalid"
const ftdIncomAttrErr1 = "Incoming arguments are invalid %s - \"%v\""
const ftdIncomAttrErr2 = "Incoming arguments are invalid  %s - \"%v\",  %s - \"%v\""
const ftdIncomAttrErr3 = "Incoming arguments are invalid  %s - \"%v\",  %s - \"%v\",  %s - \"%v\""
const ftdAttrShouldBeInfo = "Attributes should be %v"
const errorFoundReturn = "error found exit with return "

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)

	// calling method as a field, instruct the logger
	log.SetReportCaller(true)

}

//checkNotEmpty return false when all attributes are zeroed and true when attributes are not zero (empty)
func checkAllAttrNotEmpty(i ...interface{}) bool {
	log.Trace("Enter func")
	defer log.Traceln("Leaving func")

	cntNonEmptyValues := 0

	log.Tracef("incoming attr is %v", i)

	for cnt, v := range i {
		typ := reflect.TypeOf(v)
		log.Tracef("in loop %d value is -%v- reflect.TypeOf(value) is -%v-, reflect.Zero(type) has value -%v-", cnt, v, typ, reflect.Zero(typ))

		if reflect.ValueOf(v).Interface() != reflect.Zero(typ).Interface() {
			log.Traceln("reflect.ValueOf(v).Interface() != reflect.Zero(typ).Interface() is true")
			cntNonEmptyValues++
		}
	}
	if cntNonEmptyValues == len(i) {
		log.Traceln("return true")
		return true
	}
	log.Traceln("return false")
	return false

}

//WriteStringToFile writes incoming string to specified file, returns error
func WriteStringToFile(s string, filePath string) error {
	log := log.WithFields(log.Fields{
		"s":        s,
		"filePath": filePath,
	})
	log.Trace("Enter func")
	defer log.Trace("Leave func")

	if checkAllAttrNotEmpty(s, filePath) == false {
		log.Trace("checkAllAttrNotEmpty() returned false")
		log.Errorf(ftdIncomAttrErr2, "s", s, "filePath", filePath)
		return errors.New(fmt.Sprint(ftdIncomAttrErr2, "s", s, "filePath", filePath))
	}

	log.Trace("Create file start")
	f, err := os.Create(filePath)
	if err != nil {
		log.Error(err)
		log.Traceln(errorFoundReturn, err)
		return err
	}
	log.Trace("Create file end")

	log.Trace("Start write to file the incoming string ...")
	l, err := f.WriteString(s)
	if err != nil {
		log.Traceln(errorFoundReturn, err)
		log.Error(err)
		return err
	}
	log.Traceln("variable -l- with number of bytes has a value ", l)
	log.Trace("Write to file end")

	fmt.Println(l, "bytes written to file successfully")

	log.Traceln("close file start ...")
	err = f.Close()
	if err != nil {
		log.Traceln(errorFoundReturn, err)
		log.Error(err)
		return err
	}
	log.Traceln("close file end")

	return nil
}
