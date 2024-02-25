// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package gosql

import (
	"bytes"
	"context"
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

// Attributes:
//  - ColumnNames
type SelectStatement struct {
  ColumnNames []string `thrift:"column_names,1,required" db:"column_names" json:"column_names"`
}

func NewSelectStatement() *SelectStatement {
  return &SelectStatement{}
}


func (p *SelectStatement) GetColumnNames() []string {
  return p.ColumnNames
}
func (p *SelectStatement) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetColumnNames bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetColumnNames = true
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
  if !issetColumnNames{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ColumnNames is not set"));
  }
  return nil
}

func (p *SelectStatement)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.ColumnNames =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(ctx); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.ColumnNames = append(p.ColumnNames, _elem0)
  }
  if err := iprot.ReadListEnd(ctx); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *SelectStatement) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "SelectStatement"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *SelectStatement) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "column_names", thrift.LIST, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:column_names: ", p), err) }
  if err := oprot.WriteListBegin(ctx, thrift.STRING, len(p.ColumnNames)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.ColumnNames {
    if err := oprot.WriteString(ctx, string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteListEnd(ctx); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:column_names: ", p), err) }
  return err
}

func (p *SelectStatement) Equals(other *SelectStatement) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if len(p.ColumnNames) != len(other.ColumnNames) { return false }
  for i, _tgt := range p.ColumnNames {
    _src1 := other.ColumnNames[i]
    if _tgt != _src1 { return false }
  }
  return true
}

func (p *SelectStatement) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("SelectStatement(%+v)", *p)
}

func (p *SelectStatement) Validate() error {
  return nil
}
