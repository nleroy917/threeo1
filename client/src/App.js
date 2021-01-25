import "./App.css";
import {useState} from 'react';
import axios from 'axios';


const App = () => {

  const serverBase = window.location.href;
  const [url, setUrl] = useState('')
  const [shortURL, setShortURL] = useState('')

  const generateUrl = async () => {
    let res = await axios.get(`${serverBase}shorten`, { params: { url: url } } )
    if (res.status === 200) {
      let data = res.data
      console.log(data)
      setShortURL(data.shortUrl)
    }
  }

  return (
    <div className="layout">
      <div className="landing-container">
        <h1>Enter URL to shorten</h1>
        <input 
          className="url-input" 
          type="text"
          onChange={(e) => {setUrl(e.target.value)}}
        />
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
