import "./App.css";
import {useState} from 'react';
import axios from 'axios';
import {CopyToClipboard} from 'react-copy-to-clipboard';

const App = () => {

  const serverBase = window.location.href;
  const [url, setUrl] = useState('')
  const [shortUrl, setShortUrl] = useState('')
  const [copyBtnText, setCopyBtnText] = useState('Copy')

  const generateUrl = async () => {
    setCopyBtnText('Copy')
    setShortUrl("https://threeo.one/to/qw31rd")
    if (url === '') {
      alert('Please enter valid url')
    }
    let res = await axios.get(`${serverBase}shorten`, { params: { url: url } } )
    if (res.status === 200) {
      let data = res.data
      // console.log(data)
      setShortUrl(data.shortUrl)
    }
  }

  return (
    <div className="layout">
      <div className="landing-container">
        <h1 className="landing-title">threeo.one</h1>
        <p className="landing-body">Threeo.one (<em>"three-oh"</em> ) is a URL shortening service to turn even the longest of URL's into manageable and easily-sharable short-links. Simply enter your URL below and generate a short-link.</p>
        <input
          placeholder="Enter URL:"
          className="url-input" 
          type="text"
          onChange={(e) => {setUrl(e.target.value)}}
        />
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
                  onCopy={()=>{setCopyBtnText('Copied!')}}
                >
                  <button className="copy-btn">{`${copyBtnText}`}</button>
                </CopyToClipboard>
              </div>
            : <div></div>
          }
        </div>
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
