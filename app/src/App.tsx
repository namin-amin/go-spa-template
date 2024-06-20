import { createSignal } from 'solid-js'
import './App.css'

function App() {
  const [count, setCount] = createSignal("")

  const getdata = (text:string) => {
    fetch(`api/${text}`)
    .then(async res => {
      setCount(await res.text())
    }).catch(err=>{
      console.log(err);
      
    })
  }

  return (
    <>
      <h1>{count()}</h1>
      <input type="text" onKeyUp={(e)=>getdata((e.target as HTMLInputElement).value)} />
    </>
  )
}

export default App
