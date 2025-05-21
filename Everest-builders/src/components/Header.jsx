import Logo from '../assets/logo.png'

export default function Header(){
    return(
        <section className='head'>
            <img src={Logo} alt="Logo" className="logo"/>
            <h1>Everest Builders</h1>
        </section>
    )
}