import { Dialog, DialogTitle, DialogContent, DialogActions, Button } from '@mui/material';

// eslint-disable-next-line react/prop-types
const ErrorModal = ({ open, message, onClose }) => {
    return (
      <Dialog open={open} onClose={onClose}>
        <DialogTitle>Error Message</DialogTitle>
        <DialogContent>
          <p>{message}</p>
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose} color="primary">OK</Button>
        </DialogActions>
      </Dialog>
    );
  };

  export default ErrorModal