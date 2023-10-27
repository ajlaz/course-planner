import './App.css';
import Dashboard from './components/Dashboard';
import Login from './components/Login';
import Register from './components/Register';
import { useSelector } from 'react-redux';
import { useDispatch } from 'react-redux';
import { toggleLogin } from './features/appState/appState';
function App() {
  const user = useSelector(state => state.users.user);
  const login = useSelector(state => state.appState.appState.login)
  const dispatch = useDispatch();
  return (
    <div className="App">{
      user.id ? (
        <Dashboard />
      ) : (
        login ? (<div>
          <Login />
          <button onClick={() => dispatch(toggleLogin())}>Register</button>
        </div>) : (
          <div>
          <Register />
          <button onClick={() => dispatch(toggleLogin())}>Login</button>
          </div>
          
        )
      )
    }
    </div>
  );
}

export default App;
