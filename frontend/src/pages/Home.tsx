import React from 'react'
import Card from '../components/Card/Card';

function Home() {
    return <>
        home page
        <br />
        {[1, 2, 3, 4, 5].map(
            i => <Card name={`card #${i}`} number={i} />
        )}
    </>;
}

export default Home;