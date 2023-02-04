import React, {useState, useEffect} from 'react'
import '../Styles/navbar.scss'
import {Link, withRouter, useLocation} from 'react-router-dom'

export default function NavBar(props){
	const [current, setCurrent] = useState(0);	
	const location = useLocation()

	useEffect(_=>{
		if(location.pathname.includes("/user/home"))	setCurrent(0);
		if(location.pathname.includes("/user/results"))	setCurrent(1);
		if(location.pathname.includes("/vendor/home"))	setCurrent(2);
	},[location])

	return (
		<div className = {props.className}>
			<div className="navigation">
			  <ul>
				<li className={`list ${current==0?"active":""}`} data-color="#dc143c" onClick={_=>setCurrent(0)}>
				  <Link to="/user/home">
					<span className="icon"><i className={`${current==0?"fas":"far"} fa-home`}></i></span>
				  </Link>
				</li>
				<li className={`list ${current==1?"active":""}`} data-color="#3c40c6" onClick={_=>setCurrent(1)}>
				  <Link to="/user/results/nearby">
					<span className="icon"><i className={`${current==1?"fas":"far"} fa-map`}></i></span>
				  </Link>
				</li>
				<li className={`list ${current==2?"active":""}`} data-color="#05c46b" onClick={_=>setCurrent(2)}>
				  <Link to="/vendor/home">
					<span className="icon"><i className={`${current==2?"fas":"far"} fa-cog`}></i></span>
				  </Link>
				</li>
				<div className="indicator"></div>
			  </ul>
			</div>
		</div>
	)
}
