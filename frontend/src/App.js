import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import EntryList from './components/EntryList';
import EntryForm from './components/EntryForm';
import EntryDetails from './components/EntryDetails';
import NavigationBar from './components/Navbar';

function App() {
  return (
    <Router>
      <NavigationBar />
      <Routes>
        <Route path='/' element={<EntryList />} />
        <Route path='/add' element={<EntryForm />} />
        <Route path='/edit/:id' element={<EntryForm />} />
        <Route path='/entry/:id' element={<EntryDetails />} />
      </Routes>
    </Router>
  );
}

export default App;
