import { useState } from 'react'
import HomePage from './components/Home'
import AboutUs from './components/AboutUs'
import Careers from './components/Career/Career'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <HomePage/>
      <AboutUs/>
      <Careers/>
    </>
  )
}

export default App
