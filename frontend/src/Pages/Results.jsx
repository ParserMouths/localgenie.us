import React, {useEffect, useRef, useState} from 'react'
import { MapContainer, TileLayer, useMap, Marker, Popup, CircleMarker } from 'react-leaflet'
import ResultsList from '../Components/ResultsList.jsx'
import VendorLanding from '../Pages/VendorLanding.jsx'
import {sbQueryAroundUser} from '../utils/storyblok.js'
import {
  TransitionGroup,
  CSSTransition
} from "react-transition-group";
import {useParams,useRouteMatch, Route, Switch, useLocation} from "react-router-dom";

import "../Styles/results.scss"

const dummyData = [
	{id: '982hadf', imgSrc: require("../Assets/vendor-1.png"), title:"Gavin Belson's Sandwich", description: "After working in tech, i finally decided to sell the sandwiches.", cords:[23.022607, 72.5712343]},
	{id: 'd8f94fj', imgSrc: require("../Assets/vendor-2.png"), title:"Russ Hennman's Samosas", description: "After working in tech, i finally decided to sell the Fruits.", cords:[23.024453, 72.5712619]},
	{id: '89204jf', imgSrc: require("../Assets/vendor-3.png"), title:"Monica's Fruit Bowl", description: "After working in tech, i finally decided to sell the Samosas.", cords:[23.020506, 72.5713370]},
]

function ChangeView({ center }) {
  const map = useMap();
  map.setView(center, map.getZoom());
  return null;
}

function useQuery() {
  const { search } = useLocation();

  return React.useMemo(() => new URLSearchParams(search), [search]);
}

export default function Results(props){
	const removeAttribution = _ => document.querySelector('.leaflet-bottom.leaflet-right').remove()
	const [currentMarker, setCurrentMarker] = useState({Latitude: parseFloat(localStorage.getItem('latitude')), Longitude: parseFloat(localStorage.getItem('longitude'))})
	const query = useQuery()
	let {searchQuery} = useParams()
	const [data, setData] = useState([])
	const [loading, setLoading] = useState(true)
	useEffect(_=>{
		(async ()=>{
			const searchterm = searchQuery=='nearby'? '':searchQuery
			const data = await sbQueryAroundUser(searchterm)
			setData(data)
			setLoading(false)
			console.log(data)
		})()
	},[])

	// console.log("id, query", id, query)
	useEffect(_=>{
		console.log("in useeffect", searchQuery)
	}, [])
	//useEffect(_=>{
		//for(let e in dummyData)	refs.push(useRef(null))
	//},[])

	return (
		<>
			<MapContainer center={[currentMarker.Latitude, currentMarker.Longitude]} zoom={15} className="map-container" whenReady={removeAttribution}>
			  	<ChangeView center={[currentMarker.Latitude, currentMarker.Longitude]} />				
			  <TileLayer
				url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
			  />
				{
					loading ?
					<p> Loading... </p>:
					data.map((d, i)=>(
						<CircleMarker key={i} center={[d['Latitude'], d['Longitude']]} radius={8} pathOptions = {{color:'red'}}>
							<Popup>
								{d['StallName']}
							</Popup>
						  </CircleMarker>
						))}
				<Marker position={[currentMarker.Latitude, currentMarker.Longitude]}>
					<Popup>
						{currentMarker['StallName']}
					</Popup>
			  	</Marker>
			</MapContainer>
			<ResultsList data={data} className="results-list" setCurrentMarker={setCurrentMarker}/>
			<TransitionGroup>
				  <CSSTransition
					key={query}
					classNames="slide"
					timeout={300}
				  >
					{/* sneaky workaround to prevent parent to re render */}
					{query.get("StallID") ? <VendorLanding className='vendor-landing-home'/> : <i/>}
				  </CSSTransition>
			</TransitionGroup>

		</>
	)
}
