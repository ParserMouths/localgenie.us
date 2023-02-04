import React, { useEffect, useState } from 'react'
import VendorCard from './VendorCard.jsx'

import "../Styles/vendorlist.scss"
import {sbQueryAroundUser} from '../utils/storyblok.js'

// const dummyData = [
// 	{id:'ojaifj20', imgSrc: require("../Assets/vendor-1.png"), title:"Gavin Belson's Sandwich", description: "After working in tech, i finally decided to sell the sandwiches.", cords:[23.022607, 72.5712343],},
// 	{id:'oaldfadg', imgSrc: require("../Assets/vendor-2.png"), title:"Russ Hennman's Samosas", description: "After working in tech, i finally decided to sell the Fruits.", cords:[23.024453, 72.5712619]},
// 	{id:'anjaskad', imgSrc: require("../Assets/vendor-3.png"), title:"Monica's Fruit Bowl", description: "After working in tech, i finally decided to sell the Samosas.", cords:[23.020506, 72.5713370]},
// ]

							//to = {`/user/home/${d['id']}`}
export default function VendorList(props){
	const [data, setData] = useState([])
	const [loading, setLoading] = useState(true)
	useEffect(_=>{
		(async ()=>{
			const data = await sbQueryAroundUser('')
			setData(data)
			setLoading(false)
			console.log(data)
		})()
	},[])
	return (
		<div className = {props.className}>
			<div className = 'vendor-list-wrapper'>
				<h4> {props.title} </h4>
				{
					loading?
					<p>loading...</p>:
					data.map(
						(d, i)=>(
							<VendorCard 
								key={i} 
								className="vendor-card" 
								data={d} 
								setCurrentMarker={props.setCurrentMarker}
								to = {{search:`StallID=${d['StallID']}`}}
							/>
						)
				)}
			</div>
		</div>
	)
}
