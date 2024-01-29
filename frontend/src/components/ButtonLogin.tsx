import { Button } from './ui/button';
import LoginModal from './LoginModal';
import { useModal } from '../hooks/useModal';

const ButtonLogin = () => {
	const modal = useModal();

	const onLoginClick = () => {
		modal.openModal(<LoginModal />);
	};

	return (
		<Button
			onClick={onLoginClick}
			className='absolute top-8 right-10'>
			Login
		</Button>
	);
};

export default ButtonLogin;
