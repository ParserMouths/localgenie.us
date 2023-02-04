import React from 'react'
import '../Styles/button.scss'

export default function MyButton(props){
	return (
		<button className={`${props.outlined ? 'btn-outlined' : 'btn'} ${props.className}`} onClick={props.onClick}>
			{props.children}
		</button>
	)
}
