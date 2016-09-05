// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package loggingclient

import (
	"bytes"
	"fmt"
	"git.corp.plu.cn/plugo/infrastructure/loggingclient/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type TLogEntity struct {
	Title   string            `thrift:"Title,1" json:"Title"`
	Message string            `thrift:"Message,2" json:"Message"`
	Level   int8              `thrift:"Level,3" json:"Level"`
	Time    int64             `thrift:"Time,4" json:"Time"`
	Source  string            `thrift:"Source,5" json:"Source"`
	Thread  int32             `thrift:"Thread,6" json:"Thread"`
	Tags    map[string]string `thrift:"Tags,7" json:"Tags"`
}

func NewTLogEntity() *TLogEntity {
	return &TLogEntity{}
}

func (p *TLogEntity) GetTitle() string {
	return p.Title
}

func (p *TLogEntity) GetMessage() string {
	return p.Message
}

func (p *TLogEntity) GetLevel() int8 {
	return p.Level
}

func (p *TLogEntity) GetTime() int64 {
	return p.Time
}

func (p *TLogEntity) GetSource() string {
	return p.Source
}

func (p *TLogEntity) GetThread() int32 {
	return p.Thread
}

func (p *TLogEntity) GetTags() map[string]string {
	return p.Tags
}
func (p *TLogEntity) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.ReadField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *TLogEntity) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Title = v
	}
	return nil
}

func (p *TLogEntity) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *TLogEntity) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadByte(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		temp := int8(v)
		p.Level = temp
	}
	return nil
}

func (p *TLogEntity) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Time = v
	}
	return nil
}

func (p *TLogEntity) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.Source = v
	}
	return nil
}

func (p *TLogEntity) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 6: %s", err)
	} else {
		p.Thread = v
	}
	return nil
}

func (p *TLogEntity) ReadField7(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.Tags = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val1 = v
		}
		p.Tags[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *TLogEntity) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TLogEntity"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *TLogEntity) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Title", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:Title: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Title)); err != nil {
		return fmt.Errorf("%T.Title (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:Title: %s", p, err)
	}
	return err
}

func (p *TLogEntity) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Message", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:Message: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return fmt.Errorf("%T.Message (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:Message: %s", p, err)
	}
	return err
}

func (p *TLogEntity) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Level", thrift.BYTE, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:Level: %s", p, err)
	}
	if err := oprot.WriteByte(byte(p.Level)); err != nil {
		return fmt.Errorf("%T.Level (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:Level: %s", p, err)
	}
	return err
}

func (p *TLogEntity) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Time", thrift.I64, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:Time: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Time)); err != nil {
		return fmt.Errorf("%T.Time (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:Time: %s", p, err)
	}
	return err
}

func (p *TLogEntity) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Source", thrift.STRING, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:Source: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Source)); err != nil {
		return fmt.Errorf("%T.Source (5) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:Source: %s", p, err)
	}
	return err
}

func (p *TLogEntity) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Thread", thrift.I32, 6); err != nil {
		return fmt.Errorf("%T write field begin error 6:Thread: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Thread)); err != nil {
		return fmt.Errorf("%T.Thread (6) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 6:Thread: %s", p, err)
	}
	return err
}

func (p *TLogEntity) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Tags", thrift.MAP, 7); err != nil {
		return fmt.Errorf("%T write field begin error 7:Tags: %s", p, err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Tags)); err != nil {
		return fmt.Errorf("error writing map begin: %s", err)
	}
	for k, v := range p.Tags {
		if err := oprot.WriteString(string(k)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return fmt.Errorf("error writing map end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 7:Tags: %s", p, err)
	}
	return err
}

func (p *TLogEntity) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TLogEntity(%+v)", *p)
}

type TMetricEntity struct {
	Name  string            `thrift:"Name,1" json:"Name"`
	Value float64           `thrift:"Value,2" json:"Value"`
	Time  int64             `thrift:"Time,3" json:"Time"`
	Tags  map[string]string `thrift:"Tags,4" json:"Tags"`
}

func NewTMetricEntity() *TMetricEntity {
	return &TMetricEntity{}
}

func (p *TMetricEntity) GetName() string {
	return p.Name
}

func (p *TMetricEntity) GetValue() float64 {
	return p.Value
}

func (p *TMetricEntity) GetTime() int64 {
	return p.Time
}

func (p *TMetricEntity) GetTags() map[string]string {
	return p.Tags
}
func (p *TMetricEntity) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *TMetricEntity) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *TMetricEntity) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *TMetricEntity) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Time = v
	}
	return nil
}

func (p *TMetricEntity) ReadField4(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]string, size)
	p.Tags = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key2 = v
		}
		var _val3 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_val3 = v
		}
		p.Tags[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *TMetricEntity) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TMetricEntity"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *TMetricEntity) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Name", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:Name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return fmt.Errorf("%T.Name (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:Name: %s", p, err)
	}
	return err
}

func (p *TMetricEntity) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Value", thrift.DOUBLE, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:Value: %s", p, err)
	}
	if err := oprot.WriteDouble(float64(p.Value)); err != nil {
		return fmt.Errorf("%T.Value (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:Value: %s", p, err)
	}
	return err
}

func (p *TMetricEntity) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Time", thrift.I64, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:Time: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Time)); err != nil {
		return fmt.Errorf("%T.Time (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:Time: %s", p, err)
	}
	return err
}

func (p *TMetricEntity) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("Tags", thrift.MAP, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:Tags: %s", p, err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Tags)); err != nil {
		return fmt.Errorf("error writing map begin: %s", err)
	}
	for k, v := range p.Tags {
		if err := oprot.WriteString(string(k)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return fmt.Errorf("error writing map end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:Tags: %s", p, err)
	}
	return err
}

func (p *TMetricEntity) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TMetricEntity(%+v)", *p)
}

type TLogPackage struct {
	IP          int64            `thrift:"IP,1" json:"IP"`
	AppId       int32            `thrift:"AppId,2" json:"AppId"`
	LogItems    []*TLogEntity    `thrift:"LogItems,3" json:"LogItems"`
	MetricItems []*TMetricEntity `thrift:"MetricItems,4" json:"MetricItems"`
}

func NewTLogPackage() *TLogPackage {
	return &TLogPackage{}
}

func (p *TLogPackage) GetIP() int64 {
	return p.IP
}

func (p *TLogPackage) GetAppId() int32 {
	return p.AppId
}

func (p *TLogPackage) GetLogItems() []*TLogEntity {
	return p.LogItems
}

func (p *TLogPackage) GetMetricItems() []*TMetricEntity {
	return p.MetricItems
}
func (p *TLogPackage) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *TLogPackage) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *TLogPackage) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.AppId = v
	}
	return nil
}

func (p *TLogPackage) ReadField3(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*TLogEntity, 0, size)
	p.LogItems = tSlice
	for i := 0; i < size; i++ {
		_elem4 := &TLogEntity{}
		if err := _elem4.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem4, err)
		}
		p.LogItems = append(p.LogItems, _elem4)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *TLogPackage) ReadField4(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*TMetricEntity, 0, size)
	p.MetricItems = tSlice
	for i := 0; i < size; i++ {
		_elem5 := &TMetricEntity{}
		if err := _elem5.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem5, err)
		}
		p.MetricItems = append(p.MetricItems, _elem5)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *TLogPackage) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TLogPackage"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *TLogPackage) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("IP", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:IP: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.IP)); err != nil {
		return fmt.Errorf("%T.IP (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:IP: %s", p, err)
	}
	return err
}

func (p *TLogPackage) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("AppId", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:AppId: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.AppId)); err != nil {
		return fmt.Errorf("%T.AppId (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:AppId: %s", p, err)
	}
	return err
}

func (p *TLogPackage) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("LogItems", thrift.LIST, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:LogItems: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.LogItems)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.LogItems {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:LogItems: %s", p, err)
	}
	return err
}

func (p *TLogPackage) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("MetricItems", thrift.LIST, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:MetricItems: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.MetricItems)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.MetricItems {
		if err := v.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", v, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:MetricItems: %s", p, err)
	}
	return err
}

func (p *TLogPackage) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TLogPackage(%+v)", *p)
}