package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	cid "github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/encoding/dagcbor"
	ipldfluent "github.com/ipld/go-ipld-prime/fluent"
	ipldfree "github.com/ipld/go-ipld-prime/impl/free"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	multihash "github.com/multiformats/go-multihash"

	govaluate "github.com/Knetic/govaluate"
)

func blockBuilder(node ipld.Node) (ipld.Link, []byte, error) {
	linkBuilder := cidlink.LinkBuilder{cid.Prefix{
		Version:  1,
		Codec:    cid.DagCBOR,
		MhType:   multihash.SHA2_256,
		MhLength: -1,
	}}

	buf := bytes.Buffer{}
	lnk, err := linkBuilder.Build(context.Background(), ipld.LinkContext{}, node,
		func(ipld.LinkContext) (io.Writer, ipld.StoreCommitter, error) {
			return &buf, func(lnk ipld.Link) error { return nil }, nil
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return lnk, buf.Bytes(), nil
}

func blockDecoder(buf []byte) (ipld.Node, error) {
	return dagcbor.Decoder(ipldfree.NodeBuilder(), bytes.NewReader(buf))
}

func run(f float64) error {
	var fnb = ipldfluent.WrapNodeBuilder(ipldfree.NodeBuilder())
	lnk, buf, err := blockBuilder(fnb.CreateFloat(f))
	if err != nil {
		return err
	}
	fmt.Printf("%v,%v,%v", f, lnk, hex.EncodeToString(buf))

	n, err := blockDecoder(buf)
	if err != nil {
		return err
	}

	f2, err := n.AsFloat()
	if err != nil {
		return err
	}

	fmt.Printf(",%v\n", f2)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("supply floats, or an expression")
		os.Exit(1)
	}

	fmt.Println("in,cid,bytes,out")

	for _, expr := range os.Args[1:] {
		expression, err := govaluate.NewEvaluableExpression(expr)
		result, err := expression.Evaluate(nil)
		if err != nil {
			panic(err)
		}
		f, ok := result.(float64)
		if !ok {
			panic(fmt.Errorf("'%v' is not a float", expr))
		}
		run(f)
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	_ = dagcbor.Encoder // force registration of dagcbor with the cidlink encoder
}
