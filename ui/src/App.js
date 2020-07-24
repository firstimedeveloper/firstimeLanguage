import React, { Fragment, useState, useEffect } from 'react';
import axios from 'axios';
//import logo from './logo.svg';
import './App.css';
import { findByTestId } from '@testing-library/react';

const useDataApi = (initialUrl = "") => {
  const [data, setData] = useState({track: [], lines: []});
  const [url, setUrl] = useState(initialUrl);

  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
 
  useEffect(() => {
    if (url !== "") {
      console.log(url)
      const fetchData = async () => {
        setIsError(false);
        setIsLoading(true);
        try {
          const result = await axios(url);
    
          setData(result.data);
        } catch (error) {
          setIsError(true);
        }
        setIsLoading(false);
      };
   
      fetchData();
    } 
  }, [url]);

  return [{data, isLoading, isError}, setUrl];
}

const makeUrl = (action="list", id, lang, tlang) => {
  let url = `https://junhyukhan.herokuapp.com/${action}}?id=${id}`
  if (id === "") {
    return ""
  }
  if (action === "new") {
    if (lang === "") {
      return ""
    }
    url += `&lang=${lang}&tlang=${tlang}`
  }
  return url
}

function Search(props) {
  const [id, setId] = useState('');
  const [lang, setLang] = useState('');
  const [tlang, setTlang] = useState('');

  // for fetching langList
  const [{ data, isLoading, isError }, doFetch] = useDataApi();

  let defaultLang = ""
  const langListView = data ? (
    data.track.map((line,idx) => {
      if (idx === 0) {
        defaultLang = line.langCode
      }
      return <option key={line.langCode} value={line.langCode}>{line.langCode}</option>
    })
  ) : (
    <option value="">Unavailable</option>
  );

  useEffect(() => {
    try {setLang(data.track[0].langCode)} catch {}
  }, [data])
  
  const handleSubmit = (e) => {
    doFetch(`https://junhyukhan.herokuapp.com/list?id=${id}`);
    if (lang !== "") {   
      props.doFetch(`https://junhyukhan.herokuapp.com/new?id=${id}&lang=${lang}&tlang=${tlang}`);
    } else {

    }
    e.preventDefault();
  }

  return ( 
    <form onSubmit={
    handleSubmit
    }> 
    <label htmlFor="id">Video ID </label>
    <input
      name="id"
      id="id"
      type="text"
      value={id}
      onChange={event => setId(event.target.value)}
    />
    <label htmlFor="lang">Available subtitles: </label>
    <select name="lang" id="lang" onChange={event => setLang(event.target.value)} data-placeholder="Choose a Language...">
      {/* <option value="">None</option> */}
      {(langListView.length === 0) ? 
      <option>None</option> : langListView}
    </select>
  <button type="submit">Search</button>
</form>
)
    
}

function App() {
  const [{ data, isLoading, isError }, doFetch] = useDataApi(
    // 'https://junhyukhan.herokuapp.com/new',
  );
  const search = <Search doFetch={doFetch}/>;

  const transcriptView = isLoading ? (
    <div>Loading...</div>
  ) : (
    <table>
      <tbody>
      {data.lines.map(line => (
        <tr key={line.start}>
          <td className="time_subtitle">{line.start}-{line.end}</td>
          <td>{line.text}</td>
        </tr>
      ))}
      </tbody>
    </table>
  );

  return (
    <div className="wrapper">
      <Fragment>
        {search}
        {isError ? <div>Something went wrong...</div> : transcriptView}
        
          
      </Fragment>
    </div>
  );
}

const langListd = {"track":[{"langCode":"de"},{"langCode":"ja"},{"langCode":"en"}]};

export default App;

//dL5oGKNlR6I
//iXLXLCcGINw
/*
<form onSubmit={event => {
      props.doFetch(`https://junhyukhan.herokuapp.com/new?id=${id}&lang=${lang}&tlang=${tlang}`);
      event.preventDefault();
  }}>
    <label htmlFor="id">Video ID </label>
    <input
      name="id"
      id="id"
      type="text"
      value={id}
      onChange={event => setId(event.target.value)}
    />
    <label htmlFor="lang">Available subtitle language </label>
    <input
      name="lang"
      id="lang"
      type="text"
      value={lang}
      onChange={event => setLang(event.target.value)}
    />
    <label for="tlang">Translate to </label>
    <select name="tlang" id="tlang" onChange={event => setTlang(event.target.value)} data-placeholder="Choose a Language...">
      <option value="">None</option>
      <option value="AF">Afrikaans</option>
      <option value="SQ">Albanian</option>
      <option value="AR">Arabic</option>
      <option value="HY">Armenian</option>
      <option value="EU">Basque</option>
      <option value="BN">Bengali</option>
      <option value="BG">Bulgarian</option>
      <option value="CA">Catalan</option>
      <option value="KM">Cambodian</option>
      <option value="ZH">Chinese (Mandarin)</option>
      <option value="HR">Croatian</option>
      <option value="CS">Czech</option>
      <option value="DA">Danish</option>
      <option value="NL">Dutch</option>
      <option value="EN">English</option>
      <option value="ET">Estonian</option>
      <option value="FJ">Fiji</option>
      <option value="FI">Finnish</option>
      <option value="FR">French</option>
      <option value="KA">Georgian</option>
      <option value="DE">German</option>
      <option value="EL">Greek</option>
      <option value="GU">Gujarati</option>
      <option value="HE">Hebrew</option>
      <option value="HI">Hindi</option>
      <option value="HU">Hungarian</option>
      <option value="IS">Icelandic</option>
      <option value="ID">Indonesian</option>
      <option value="GA">Irish</option>
      <option value="IT">Italian</option>
      <option value="JA">Japanese</option>
      <option value="JW">Javanese</option>
      <option value="KO">Korean</option>
      <option value="LA">Latin</option>
      <option value="LV">Latvian</option>
      <option value="LT">Lithuanian</option>
      <option value="MK">Macedonian</option>
      <option value="MS">Malay</option>
      <option value="ML">Malayalam</option>
      <option value="MT">Maltese</option>
      <option value="MI">Maori</option>
      <option value="MR">Marathi</option>
      <option value="MN">Mongolian</option>
      <option value="NE">Nepali</option>
      <option value="NO">Norwegian</option>
      <option value="FA">Persian</option>
      <option value="PL">Polish</option>
      <option value="PT">Portuguese</option>
      <option value="PA">Punjabi</option>
      <option value="QU">Quechua</option>
      <option value="RO">Romanian</option>
      <option value="RU">Russian</option>
      <option value="SM">Samoan</option>
      <option value="SR">Serbian</option>
      <option value="SK">Slovak</option>
      <option value="SL">Slovenian</option>
      <option value="ES">Spanish</option>
      <option value="SW">Swahili</option>
      <option value="SV">Swedish </option>
      <option value="TA">Tamil</option>
      <option value="TT">Tatar</option>
      <option value="TE">Telugu</option>
      <option value="TH">Thai</option>
      <option value="BO">Tibetan</option>
      <option value="TO">Tonga</option>
      <option value="TR">Turkish</option>
      <option value="UK">Ukrainian</option>
      <option value="UR">Urdu</option>
      <option value="UZ">Uzbek</option>
      <option value="VI">Vietnamese</option>
      <option value="CY">Welsh</option>
      <option value="XH">Xhosa</option>
    </select>
    <button type="submit">Search</button>
  </form>
  );
*/