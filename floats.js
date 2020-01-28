const Block = require('@ipld/block')
if (process.argv.length < 3) throw new Error('supply some floats, or expressions')
async function run () {
  console.log('in,cid,bytes,out')
  for (let i = 2; i < process.argv.length; i++) {
    const f = eval(process.argv[i])
    const b = Block.encoder(f, 'dag-cbor')
    const cid = await b.cid()
    const b2 = Block.create(b.encode(), cid)
    console.log(`${f},${cid},${b.encode().toString('hex')},${b2.decode()}`)
  }
}

run().catch((err) => {
  console.log(err)
  process.exit(1)
})
