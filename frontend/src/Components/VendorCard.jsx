import React, {useEffect} from 'react'
import {Link} from 'react-router-dom'
import { useInView } from 'react-intersection-observer';
import '../Styles/vendorcard.scss'

export default function VendorCard(props){
	const { ref, inView, entry } = useInView({
		threshold: 0.65,
	});
	useEffect(_=>{
		if (inView){
			if (props['setCurrentMarker'])	props.setCurrentMarker(props.data)
		}
	}, [inView])
	return (
		<div className = {props.className} ref = {ref}>
		<Link to = {props.to} style={{textDecoration:'None'}}>
			<div className={`vendor-wrapper ${props.animate? inView?'grow':'shrink':''}`} style={{width: props.width}}>
				<img src={props.data['Assets'].length>0 && props.data['Assets'][0]}/>
				<div className='title-description'>
					<h3> {props.data['StallName']} </h3>
					<p> {props.data['Offering']} </p>
				</div>
			</div>
		</Link>
		</div>
	)
}
