import { useState } from 'react'
import HomePage from './components/Home'
import AboutUs from './components/AboutUs'
import Careers from './components/Career/Career'
import Login from './components/Login'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <HomePage/>
      <AboutUs/>
      <Careers/>
      <Login/>
    </>
  )
}

export default App
