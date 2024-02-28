// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package exception

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"strings"
	"regexp"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal
// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

type ExceptionCode int64
const (
  ExceptionCode_INVALID_SYNTAX ExceptionCode = 1
)

func (p ExceptionCode) String() string {
  switch p {
  case ExceptionCode_INVALID_SYNTAX: return "INVALID_SYNTAX"
  }
  return "<UNSET>"
}

func ExceptionCodeFromString(s string) (ExceptionCode, error) {
  switch s {
  case "INVALID_SYNTAX": return ExceptionCode_INVALID_SYNTAX, nil 
  }
  return ExceptionCode(0), fmt.Errorf("not a valid ExceptionCode string")
}


func ExceptionCodePtr(v ExceptionCode) *ExceptionCode { return &v }

func (p ExceptionCode) MarshalText() ([]byte, error) {
return []byte(p.String()), nil
}

func (p *ExceptionCode) UnmarshalText(text []byte) error {
q, err := ExceptionCodeFromString(string(text))
if (err != nil) {
return err
}
*p = q
return nil
}

func (p *ExceptionCode) Scan(value interface{}) error {
v, ok := value.(int64)
if !ok {
return errors.New("Scan value is not int64")
}
*p = ExceptionCode(v)
return nil
}

func (p * ExceptionCode) Value() (driver.Value, error) {
  if p == nil {
    return nil, nil
  }
return int64(*p), nil
}
// Attributes:
//  - Code
//  - Description
type InvalidSyntaxException struct {
  Code ExceptionCode `thrift:"code,1,required" db:"code" json:"code"`
  Description string `thrift:"description,2,required" db:"description" json:"description"`
}

func NewInvalidSyntaxException() *InvalidSyntaxException {
  return &InvalidSyntaxException{
Code: 1,
}
}


func (p *InvalidSyntaxException) GetCode() ExceptionCode {
  return p.Code
}

func (p *InvalidSyntaxException) GetDescription() string {
  return p.Description
}
func (p *InvalidSyntaxException) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetCode bool = false;
  var issetDescription bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetCode = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetDescription = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetCode{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Code is not set"));
  }
  if !issetDescription{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Description is not set"));
  }
  return nil
}

func (p *InvalidSyntaxException)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := ExceptionCode(v)
  p.Code = temp
}
  return nil
}

func (p *InvalidSyntaxException)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Description = v
}
  return nil
}

func (p *InvalidSyntaxException) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "InvalidSyntaxException"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *InvalidSyntaxException) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "code", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.Code)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err) }
  return err
}

func (p *InvalidSyntaxException) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "description", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:description: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Description)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.description (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:description: ", p), err) }
  return err
}

func (p *InvalidSyntaxException) Equals(other *InvalidSyntaxException) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.Code != other.Code { return false }
  if p.Description != other.Description { return false }
  return true
}

func (p *InvalidSyntaxException) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("InvalidSyntaxException(%+v)", *p)
}

func (p *InvalidSyntaxException) Error() string {
  return p.String()
}

func (InvalidSyntaxException) TExceptionType() thrift.TExceptionType {
  return thrift.TExceptionTypeCompiled
}

var _ thrift.TException = (*InvalidSyntaxException)(nil)

func (p *InvalidSyntaxException) Validate() error {
  return nil
}