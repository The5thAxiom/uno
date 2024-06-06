import {BrowserRouter, Routes, Route} from 'react-router-dom'
import Home from './pages/Home'
import Footer from './components/Footer'

function App() {
  return (
    <>
      <header>
        UNO
        <nav>links</nav>
      </header>
      <main>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<Home/>}/>
            <Route path='/play' element={<>play</>}/>
          </Routes>
        </BrowserRouter>
      </main>
      <Footer />
    </>
  )
}

export default App
