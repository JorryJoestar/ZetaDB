package main

import (
	"ZetaDB/container"
	. "ZetaDB/execution/querySubOperator"
	"ZetaDB/parser"
	. "ZetaDB/storage"
	. "ZetaDB/utility"
	"fmt"
	"sync"
)

type Kernel struct {
	parser *parser.Parser
}

//for singleton pattern
var instance *Kernel
var once sync.Once

//to get kernel, call this function
func GetInstance() *Kernel {
	once.Do(func() {
		instance = &Kernel{
			parser: parser.GetInstance()}
	})
	return instance
}

func main() {
	kernel := GetInstance()

	s := "select a,b,c from b;"

	ast := kernel.parser.ParseSql(s)
	fmt.Println(ASTToString(ast))

	se := GetStorageEngine(DEFAULT_DATAFILE_LOCATION, DEFAULT_INDEXFILE_LOCATION, DEFAULT_LOGFILE_LOCATION)

	domain1, _ := container.NewDomain("id", container.INT, 0, 0)
	domain2, _ := container.NewDomain("name", container.VARCHAR, 20, 0)
	domain3, _ := container.NewDomain("address", container.VARCHAR, 5000, 0)
	domain4, _ := container.NewDomain("birthdate", container.DATE, 0, 0)
	var domainList []*container.Domain
	domainList = append(domainList, domain1)
	domainList = append(domainList, domain2)
	domainList = append(domainList, domain3)
	domainList = append(domainList, domain4)
	schema, _ := container.NewSchema("testTable", domainList, nil)

	/*

		t1field0, _ := container.NewFieldFromBytes(INTToBytes(1))
		t1Name, _ := VARCHARToBytes("simeon")
		t1field1, _ := container.NewFieldFromBytes(t1Name)
		t1address, _ := VARCHARToBytes("Birmingham")
		t1field2, _ := container.NewFieldFromBytes(t1address)
		t1birthdate, _ := DATEToBytes("1997-11-12")
		t1field3, _ := container.NewFieldFromBytes(t1birthdate)
		var t1fields []*container.Field
		t1fields = append(t1fields, t1field0)
		t1fields = append(t1fields, t1field1)
		t1fields = append(t1fields, t1field2)
		t1fields = append(t1fields, t1field3)
		tuple1, _ := container.NewTuple(30, 1, schema, t1fields)

		t2field0, _ := container.NewFieldFromBytes(INTToBytes(2))
		t2Name, _ := VARCHARToBytes("Alex")
		t2field1, _ := container.NewFieldFromBytes(t2Name)
		t2address, _ := VARCHARToBytes("Beijing")
		t2field2, _ := container.NewFieldFromBytes(t2address)
		t2birthdate, _ := DATEToBytes("1998-03-02")
		t2field3, _ := container.NewFieldFromBytes(t2birthdate)
		var t2fields []*container.Field
		t2fields = append(t2fields, t2field0)
		t2fields = append(t2fields, t2field1)
		t2fields = append(t2fields, t2field2)
		t2fields = append(t2fields, t2field3)
		tuple2, _ := container.NewTuple(30, 2, schema, t2fields)

		t3field0, _ := container.NewFieldFromBytes(INTToBytes(3))
		t3Name, _ := VARCHARToBytes("Claire")
		t3field1, _ := container.NewFieldFromBytes(t3Name)
		var longX string
		for i := 0; i < 5000; i++ {
			longX += "X"
		}
		t3address, _ := VARCHARToBytes(longX)
		t3field2, _ := container.NewFieldFromBytes(t3address)
		t3birthdate, _ := DATEToBytes("1997-10-20")
		t3field3, _ := container.NewFieldFromBytes(t3birthdate)
		var t3fields []*container.Field
		t3fields = append(t3fields, t3field0)
		t3fields = append(t3fields, t3field1)
		t3fields = append(t3fields, t3field2)
		t3fields = append(t3fields, t3field3)
		tuple3, _ := container.NewTuple(30, 3, schema, t3fields)

		t4field0, _ := container.NewFieldFromBytes(INTToBytes(4))
		t4Name, _ := VARCHARToBytes("Woozie")
		t4field1, _ := container.NewFieldFromBytes(t4Name)
		t4address, _ := VARCHARToBytes("Los Santos")
		t4field2, _ := container.NewFieldFromBytes(t4address)
		t4birthdate, _ := DATEToBytes("1959-10-10")
		t4field3, _ := container.NewFieldFromBytes(t4birthdate)
		var t4fields []*container.Field
		t4fields = append(t4fields, t4field0)
		t4fields = append(t4fields, t4field1)
		t4fields = append(t4fields, t4field2)
		t4fields = append(t4fields, t4field3)
		tuple4, _ := container.NewTuple(30, 4, schema, t4fields)

		t3Bytes, _ := tuple3.TupleToBytes()
		p63Data := t3Bytes[:4064]
		t3Bytes = t3Bytes[4064:]
		p243Data := t3Bytes

		p72 := NewDataPageMode0(72, 30, 72, 63)
		p72.InsertTuple(tuple1)
		p72.InsertTuple(tuple2)
		p63 := NewDataPageMode1(63, 30, 72, 19, 243, p63Data)
		p243 := NewDataPageMode2(243, 30, 963, 63, 243, p243Data)
		p19 := NewDataPageMode0(19, 30, 63, 19)
		p19.InsertTuple(tuple4)

		se.InsertDataPage(p72)
		se.InsertDataPage(p63)
		se.InsertDataPage(p243)
		se.InsertDataPage(p19)

		se.SwapDataPage(p72.DpGetPageId())
		se.SwapDataPage(p63.DpGetPageId())
		se.SwapDataPage(p243.DpGetPageId())
		se.SwapDataPage(p19.DpGetPageId()) */

	/* 	it := NewSequentialFileReaderIterator(se, 72, schema)
	   	it.Open(nil, nil)

	   	for it.HasNext() {
	   		t, _ := it.GetNext()
	   		fmt.Println("------------")
	   		idBytes, _ := t.TupleGetFieldValue(0)
	   		id, _ := BytesToINT(idBytes)
	   		nameBytes, _ := t.TupleGetFieldValue(1)
	   		name, _ := BytesToVARCHAR(nameBytes)
	   		addressBytes, _ := t.TupleGetFieldValue(2)
	   		address, _ := BytesToVARCHAR(addressBytes)
	   		birthdateBytes, _ := t.TupleGetFieldValue(3)
	   		birthdate, _ := BytesToDATE(birthdateBytes)
	   		fmt.Println()
	   		fmt.Printf("id:        %v\n", id)
	   		fmt.Printf("name:      %v\n", name)
	   		fmt.Printf("address:   %v\n", address)
	   		fmt.Printf("birthdate: %v\n", birthdate)
	   		fmt.Println()
	   		fmt.Println("------------")

	   	} */

	/* 	t0f0, _ := container.NewFieldFromBytes(INTToBytes(1))
	   	t0Name, _ := VARCHARToBytes("Simeon")
	   	t0f1, _ := container.NewFieldFromBytes(t0Name)
	   	t0Address, _ := VARCHARToBytes("Birmingham")
	   	t0f2, _ := container.NewFieldFromBytes(t0Address)
	   	t0Date, _ := DATEToBytes("1997-11-12")
	   	t0f3, _ := container.NewFieldFromBytes(t0Date)
	   	var f0 []*container.Field
	   	f0 = append(f0, t0f0)
	   	f0 = append(f0, t0f1)
	   	f0 = append(f0, t0f2)
	   	f0 = append(f0, t0f3)
	   	t0, _ := container.NewTuple(30, 0, schema, f0)

	   	t1f0, _ := container.NewFieldFromBytes(INTToBytes(2))
	   	t1Name, _ := VARCHARToBytes("Alex")
	   	t1f1, _ := container.NewFieldFromBytes(t1Name)
	   	t1Address, _ := VARCHARToBytes("Beijing")
	   	t1f2, _ := container.NewFieldFromBytes(t1Address)
	   	t1Date, _ := DATEToBytes("1998-03-02")
	   	t1f3, _ := container.NewFieldFromBytes(t1Date)
	   	var f1 []*container.Field
	   	f1 = append(f1, t1f0)
	   	f1 = append(f1, t1f1)
	   	f1 = append(f1, t1f2)
	   	f1 = append(f1, t1f3)
	   	t1, _ := container.NewTuple(30, 1, schema, f1)

	   	t2f0, _ := container.NewFieldFromBytes(INTToBytes(168))
	   	t2Name, _ := VARCHARToBytes("Claire")
	   	t2f1, _ := container.NewFieldFromBytes(t2Name)
	   	t2Address, _ := VARCHARToBytes("London")
	   	t2f2, _ := container.NewFieldFromBytes(t2Address)
	   	t2Date, _ := DATEToBytes("1997-10-20")
	   	t2f3, _ := container.NewFieldFromBytes(t2Date)
	   	var f2 []*container.Field
	   	f2 = append(f2, t2f0)
	   	f2 = append(f2, t2f1)
	   	f2 = append(f2, t2f2)
	   	f2 = append(f2, t2f3)
	   	t2, _ := container.NewTuple(30, 2, schema, f2)

	   	t3f0, _ := container.NewFieldFromBytes(INTToBytes(1))
	   	t3Name, _ := VARCHARToBytes("Simeon")
	   	t3f1, _ := container.NewFieldFromBytes(t3Name)
	   	t3Address, _ := VARCHARToBytes("Birmingham")
	   	t3f2, _ := container.NewFieldFromBytes(t3Address)
	   	t3Date, _ := DATEToBytes("1997-11-12")
	   	t3f3, _ := container.NewFieldFromBytes(t3Date)
	   	var f3 []*container.Field
	   	f3 = append(f3, t3f0)
	   	f3 = append(f3, t3f1)
	   	f3 = append(f3, t3f2)
	   	f3 = append(f3, t3f3)
	   	t3, _ := container.NewTuple(30, 3, schema, f3)

	   	t4f0, _ := container.NewFieldFromBytes(INTToBytes(72))
	   	t4Name, _ := VARCHARToBytes("Jack")
	   	t4f1, _ := container.NewFieldFromBytes(t4Name)
	   	t4Address, _ := VARCHARToBytes("Java")
	   	t4f2, _ := container.NewFieldFromBytes(t4Address)
	   	t4Date, _ := DATEToBytes("1894-11-16")
	   	t4f3, _ := container.NewFieldFromBytes(t4Date)
	   	var f4 []*container.Field
	   	f4 = append(f4, t4f0)
	   	f4 = append(f4, t4f1)
	   	f4 = append(f4, t4f2)
	   	f4 = append(f4, t4f3)
	   	t4, _ := container.NewTuple(30, 4, schema, f4)

	   	t5f0, _ := container.NewFieldFromBytes(INTToBytes(1))
	   	t5Name, _ := VARCHARToBytes("Simeon")
	   	t5f1, _ := container.NewFieldFromBytes(t5Name)
	   	t5Address, _ := VARCHARToBytes("Birmingham")
	   	t5f2, _ := container.NewFieldFromBytes(t5Address)
	   	t5Date, _ := DATEToBytes("1997-11-12")
	   	t5f3, _ := container.NewFieldFromBytes(t5Date)
	   	var f5 []*container.Field
	   	f5 = append(f5, t5f0)
	   	f5 = append(f5, t5f1)
	   	f5 = append(f5, t5f2)
	   	f5 = append(f5, t5f3)
	   	t5, _ := container.NewTuple(30, 5, schema, f5)

	   	page := NewDataPageMode0(30, 30, 30, 30)
	   	page.InsertTuple(t0)
	   	page.InsertTuple(t1)
	   	page.InsertTuple(t2)
	   	page.InsertTuple(t3)
	   	page.InsertTuple(t4)
	   	page.InsertTuple(t5)

	   	se.InsertDataPage(page)
	   	se.SwapDataPage(page.DpGetPageId()) */
	page, _ := se.GetDataPage(30, schema)
	fmt.Println(page)
	seqIt := NewSequentialFileReaderIterator(se, 30, schema)
	seqIt.Open(nil, nil)
	for seqIt.HasNext() {
		fmt.Println(seqIt.GetNext())
	}
	fmt.Println("---------------------")

	seqIt.Close()
	seqIt.Open(nil, nil)
	dupEliminateIt := NewDuplicateEliminationIterator()
	dupEliminateIt.Open(seqIt, nil)
	for dupEliminateIt.HasNext() {
		fmt.Println(dupEliminateIt.GetNext())
	}
	fmt.Println("---------------------")
	seqIt.Close()
	seqIt.Open(nil, nil)
	dupEliminateIt.Close()
	dupEliminateIt.Open(seqIt, nil)
	for dupEliminateIt.HasNext() {
		fmt.Println(dupEliminateIt.GetNext())
	}
}
