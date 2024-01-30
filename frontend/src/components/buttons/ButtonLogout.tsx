import { IconUser } from '@tabler/icons-react';
import { Button } from '../ui/button';
import { useGlobalStore } from '../../store/globalState';
import { API_URL } from '../../constants/apiUrl';

const ButtonLogout = () => {
	const { clearStore } = useGlobalStore();

	const onLogoutClick = () => {
		clearStore();
		window.location.href = `${API_URL}/auth/logout`;
	};
	return (
		<Button onClick={onLogoutClick}>
			<IconUser />
			Logout
		</Button>
	);
};

export default ButtonLogout;
