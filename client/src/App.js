import "./App.css";
import {useState} from 'react';


const App = () => {

  const [url, setUrl] = useState('')
  const [shortURL, setShortURL] = useState('')

  const generateUrl = async () => {
    setShortURL('https://threeo1/to/123456')
  }

  return (
    <div className="layout">
      <div className="landing-container">
        <h1>Enter URL to shorten</h1>
        <input className="url-input" type="text"/>
        <button
          className="generate-btn"
          onClick={()=>{generateUrl()}}
          onChange={(e)=>{setShortURL(e.target.value)}}
        >
          Generate
        </button>
        <div className="shortened-url-wrapper">
          {
            shortURL ? 
              <div>
                {`${shortURL}`}
              </div>
            : <div></div>

          }
        </div>
      </div>
    </div>
  );
}

export default App;
