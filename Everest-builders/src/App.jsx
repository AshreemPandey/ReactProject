import { useState } from 'react'
import HomePage from './components/Home'
import AboutUs from './components/AboutUs'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <HomePage/>
      <AboutUs/>
    </>
  )
}

export default App
