import { useGlobalStore } from '../store/globalState';
import ButtonLogin from './ButtonLogin';
import ButtonLogout from './ButtonLogout';

const Header = () => {
	const { user } = useGlobalStore();

	return (
		<header className='z-20 flex items-center justify-between py-5 px-10 w-full'>
			{user ? (
				<>
					<h2 className='text-[2rem] font-geistUltra'>Short - it</h2>
					<ButtonLogout />
				</>
			) : (
				<ButtonLogin />
			)}
		</header>
	);
};

export default Header;
