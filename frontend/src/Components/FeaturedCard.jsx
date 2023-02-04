import React from 'react'
import '../Styles/featuredcard.scss'

export default function FeaturedCard(props){
	return (
		<div className = {props.className}>
			<div className='wrapper'>
				<img src={props.img}/>
				<h3> {props.title} </h3>
				<p> {props.description} </p>
			</div>
		</div>
	)
}
