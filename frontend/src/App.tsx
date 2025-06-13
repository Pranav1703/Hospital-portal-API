import {BrowserRouter as Router,Routes,Route } from 'react-router-dom'
import LoginPage from './pages/Login'
import SignUpPage from './pages/Signup'
import Home from './pages/Home'

function App() {


  return (
    <>
      <Router>
        <Routes>
          <Route path='/' element={<Home/>}/>
          <Route path='/login' element={<LoginPage/>}/>
          <Route path='/signup' element={<SignUpPage/>}/>
        </Routes>
      </Router>      
    </>
  )
}

export default App
