package flow

import (
	"errors"
	"fmt"
	"strings"

	"github.com/direktiv/direktiv/pkg/flow/ent"
	derrors "github.com/direktiv/direktiv/pkg/flow/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCodeInternal               = "direktiv.internal.error"
	ErrCodeWorkflowUnparsable     = "direktiv.workflow.unparsable"
	ErrCodeMultipleErrors         = "direktiv.workflow.multipleErrors"
	ErrCodeCancelledByParent      = "direktiv.cancels.parent"
	ErrCodeSoftTimeout            = "direktiv.cancels.timeout.soft"
	ErrCodeHardTimeout            = "direktiv.cancels.timeout.hard"
	ErrCodeJQBadQuery             = "direktiv.jq.badCommand"
	ErrCodeJQNotObject            = "direktiv.jq.notObject"
	ErrCodeAllBranchesFailed      = "direktiv.parallel.allFailed"
	ErrCodeNotArray               = "direktiv.foreach.badArray"
	ErrCodeFailedSchemaValidation = "direktiv.schema.failed"
	ErrCodeJQNotString            = "direktiv.jq.notString"
	ErrCodeInvalidVariableKey     = "direktiv.var.invalidKey"
)

var (
	ErrNotDir         = errors.New("not a directory")
	ErrNotWorkflow    = errors.New("not a workflow")
	ErrNotMirror      = errors.New("not a git mirror")
	ErrMirrorLocked   = errors.New("git mirror is locked")
	ErrMirrorUnlocked = errors.New("git mirror is not locked")
)

func translateError(err error) error {

	if derrors.IsNotFound(err) {
		err = status.Error(codes.NotFound, strings.TrimPrefix(err.Error(), "ent: "))
		return err
	}

	if _, ok := err.(*ent.ConstraintError); ok {

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			err = status.Error(codes.AlreadyExists, "resource already exists")
			return err
		}

	}

	if _, ok := err.(*ent.ValidationError); ok {

		if strings.Contains(err.Error(), "validator failed") {
			err = status.Error(codes.InvalidArgument, "one or more fields has an invalid value")
			return err
		}

	}

	if strings.Contains(err.Error(), "already exists") {
		err = status.Error(codes.AlreadyExists, "resource already exists")
		return err
	}

	fmt.Println("TRANSLATE", err)

	return err

}
