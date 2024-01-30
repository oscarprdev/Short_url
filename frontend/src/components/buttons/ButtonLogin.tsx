import { Button } from '../ui/button';
import LoginModal from '../modals/LoginModal';
import { useModal } from '../../hooks/useModal';

const ButtonLogin = () => {
	const modal = useModal();

	const onLoginClick = () => {
		modal.openModal(<LoginModal />);
	};

	return (
		<Button
			data-testid='login-btn'
			onClick={onLoginClick}
			className='absolute top-8 right-10'>
			Login
		</Button>
	);
};

export default ButtonLogin;
