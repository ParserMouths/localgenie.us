import React, { useEffect, useState } from 'react'
import Logo from './logo.jsx'
import '../Styles/header.scss'
import axios from 'axios'

// window.geocoder = geocoder

export default function Header(props){
	const [city, setCity] = useState(localStorage.getItem("city_town") || "update location")
	useEffect(()=>localStorage.setItem('city_town', city),[city])
	const updateLocation = ()=>{
		const _confirm = window.confirm("Do you want to update your location?")
		if(!_confirm) return
		navigator.geolocation.getCurrentPosition(async ({coords})=>{
			console.log(coords)
			localStorage.setItem('latitude', coords.latitude)
			localStorage.setItem('longitude', coords.longitude)
			const {data} = await axios.get(`https://api.opencagedata.com/geocode/v1/json?key=c7215cdf8efb48cd8b5fd90aeb5328e1&q=${coords.latitude}+${coords.longitude}&pretty=1&no_annotations=1`)
			setCity(data.results[0].components.city || data.results[0].components.town || data.results[0].components.state_district)
		})
	}
	return (
		<div className={props.className}>
			<div className='header-wrapper'>
			  	<Logo     
					style = {{
						  margin: "0.5rem",
						  height: "2rem",
						  width: "min-content", 
						}}/>
				<span className='locator' onClick={updateLocation}>{city}<i className="far fa-location" style={{marginLeft:'5px'}}></i></span>
			</div>
		</div>
	)
}
