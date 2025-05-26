export default function List(props){ 
    return(
        <a href={props.link}><li>{props.value}</li></a>
    )
}