import { useModal } from '../../hooks/useModal';
import WrapperAction from '../containers/WrapperAction';
import LoginModal from '../modals/LoginModal';

const ButtonLogin = () => {
	const modal = useModal();

	const onLoginClick = () => {
		modal.openModal(<LoginModal />);
	};

	return (
		<div className="absolute top-5 right-5">
			<WrapperAction color="default" id="login" onClick={onLoginClick}>
				Login
			</WrapperAction>
		</div>
	);
};

export default ButtonLogin;
