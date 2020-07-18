import React, { Fragment, useState, useEffect } from 'react';
import axios from 'axios';
//import logo from './logo.svg';
import './App.css';

const useDataApi = (initialUrl) => {
  const [data, setData] = useState({lines: []});
  const [url, setUrl] = useState(initialUrl);

  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);
 
  useEffect(() => {
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
  }, [url]);

  return [{data, isLoading, isError}, setUrl];
}

function App() {
  const [id, setId] = useState('dL5oGKNlR6I');
  const [lang, setLang] = useState('de');
  const [tlang, setTlang] = useState('');

  const [{ data, isLoading, isError }, doFetch] = useDataApi(
    'https://junhyukhan.herokuapp.com/new',
  );

 
  return (
    <div class="wrapper">
      <Fragment>
        <form onSubmit={event => {
            doFetch(`https://junhyukhan.herokuapp.com/new?id=${id}&lang=${lang}&tlang=${tlang}`);
            event.preventDefault();
        }}>
          <input
            type="text"
            value={id}
            onChange={event => setId(event.target.value)}
          />
          <input
            type="text"
            value={lang}
            onChange={event => setLang(event.target.value)}
          />
          <input
            type="text"
            value={tlang}
            onChange={event => setTlang(event.target.value)}
          />
          <button type="submit">Search</button>
        </form>
        {isError && <div>Something went wrong...</div>}

        {isLoading ? (
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
        )}
          
      </Fragment>
    </div>
  );
}


export default App;
