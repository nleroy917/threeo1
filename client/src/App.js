import "./App.css";
import {useState} from 'react';
import axios from 'axios';
import {CopyToClipboard} from 'react-copy-to-clipboard';

const App = () => {

  const serverBase = window.location.href;
  const [url, setUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [copyBtnText, setCopyBtnText] = useState('Copy');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);

  const generateUrl = async () => {
    /* 
      Function to genreate a short-url. Checks to
      make sure that thre is an input. Returns
      if invalid. Then it tries to make a request
      to the server to generate one. (See server
      code for server logic). Catches and reports 
      errors, otherwise displays results
    */
    if (url === '') {
      alert('Please enter valid url')
      return
    }
    setLoading(true)
    try {
      let res = await axios.get(`${serverBase}shorten`, { params: { url: url } } )
      if (res.status === 200) {
        let data = res.data
        // console.log(data)
        setShortUrl(data.shortUrl)
        setLoading(false)
      }
    } catch {
      setError(true)
      setLoading(false)
    }
  }

  return (
    <div className="layout">
      <div className="landing-container">
      <div className="blob-top">
        <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
          <path fill="#FF5F5F" d="M23,-24.9C31.3,-7.9,40.6,1.7,39.3,9.7C38,17.6,26.2,23.8,13.2,31.6C0.2,39.4,-13.9,48.9,-31.2,47.3C-48.4,45.7,-68.6,33,-72.8,16.2C-76.9,-0.6,-65,-21.5,-50.2,-40C-35.4,-58.5,-17.7,-74.6,-5.2,-70.5C7.3,-66.4,14.7,-42,23,-24.9Z" transform="translate(100 100)" />
        </svg>
      </div>
        <h1 className="landing-title">threeo.one</h1>
        <p className="landing-body">Threeo.one (<em>"three-oh"</em> ) is a URL shortening service to turn even the longest of URL's into manageable and easily-sharable short-links. Simply enter your URL below and generate a short-link.</p>
        <div className="input-loader-wrapper">
          <input
            placeholder="Enter URL:"
            className="url-input" 
            type="text"
            onChange={(e) => {setUrl(e.target.value)}}
          />
          {loading ? <div class="lds-ring"><div></div><div></div><div></div><div></div></div> : <div></div>}
        </div>
        <button
          className="generate-btn"
          onClick={()=>{generateUrl()}}
        >
          Generate
        </button>
        <br></br>
        <div className="shortened-url-wrapper">
          {
            shortUrl ? 
              <div className="url-box-border">
                <h4 className="short-url-title">Your short-url:</h4>
                <a href={`${shortUrl}`}>{`${shortUrl}`}</a>
                <CopyToClipboard text={shortUrl}
                  onCopy={()=>{
                    setCopyBtnText('Copied!')
                    setTimeout(() => {
                      setCopyBtnText("Copy")
                    }, 1500)
                  }}
                >
                  <button className="copy-btn">{`${copyBtnText}`}</button>
                </CopyToClipboard>
              </div>
            : <div></div>
          }
        </div>
      </div>
      <div className="blob-bottom">
      <svg viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
        <path fill="#FF5F5F" d="M49.3,-69C62.5,-58.2,71,-42,74.8,-25.4C78.6,-8.9,77.8,8.1,70,19.9C62.2,31.6,47.4,38.3,34.6,46.7C21.7,55.1,10.9,65.3,-1.9,67.9C-14.7,70.5,-29.3,65.5,-39.1,56.1C-48.9,46.7,-53.9,32.9,-57.4,19.3C-60.9,5.7,-63,-7.7,-58,-17.5C-52.9,-27.2,-40.7,-33.2,-29.9,-44.8C-19,-56.3,-9.5,-73.4,4.3,-79.2C18,-85.1,36.1,-79.8,49.3,-69Z" transform="translate(100 100)" />
      </svg>
      </div>
      <footer className="footer">
          <a className="link" href="https://github.com/nleroy917/threeo1">Source</a>
          <a>v0.0.1</a>
          <a className="link">About</a>
        </footer>
    </div>
  );
}

export default App;
