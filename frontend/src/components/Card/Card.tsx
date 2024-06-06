import React from 'react'
import './Card.css'

type CardProps = {
    name: string;
    number: number;
}

function Card({name, number}: CardProps) {
    return <div className='card'>
        {number} ({name})
    </div>
}

export default Card;