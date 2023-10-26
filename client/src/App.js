import './App.css';
import Dashboard from './components/Dashboard';
import Login from './components/Login';
import { useSelector } from 'react-redux';
function App() {
  const user = useSelector(state => state.users.user);
  return (
    <div className="App">{
      user.id ? (
        <Dashboard />
      ) : <Login />
    }
    </div>
  );
}

export default App;
