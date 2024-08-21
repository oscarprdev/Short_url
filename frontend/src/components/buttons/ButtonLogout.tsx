import { API_URL } from '../../constants/apiUrl';
import { useGlobalStore } from '../../store/globalState';
import WrapperAction from '../containers/WrapperAction';

const ButtonLogout = () => {
	const { clearStore } = useGlobalStore();

	const onLogoutClick = () => {
		clearStore();
		window.location.href = `${API_URL}/auth/logout`;
	};
	return (
		<WrapperAction color="default" id="logout" onClick={onLogoutClick}>
			Logout
		</WrapperAction>
	);
};

export default ButtonLogout;
