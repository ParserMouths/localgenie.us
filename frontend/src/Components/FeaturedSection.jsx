import React from 'react'
import FeaturedCard from './FeaturedCard.jsx'

import "../Styles/featuredsection.scss"

export default function FeaturedSection(props){
	return (
		<div className = {props.className}>
			<div className = 'featured-wrapper'>
				<h4> Featured </h4>
				<div className = 'card-wrapper'>
				  	<FeaturedCard className="card" img={require('../Assets/samosa.png')} title="Samosa" description="Tasty asf" />
				  	<FeaturedCard className="card" img={require('../Assets/pani-puri.png')} title="Pani Puri" description="Spicy asf" />
				  	<FeaturedCard className="card" img={require('../Assets/fruits.png')} title="Fruits" description="Healthy asf" />
				  	<FeaturedCard className="card" img={require('../Assets/ice-cream.png')} title="Ice Cream" description="Cold asf" />
				</div>
			</div>
		</div>
	)
}
