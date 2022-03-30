use anyhow::{Result};
use sysinfo::{ProcessExt,System,SystemExt,PidExt,User,UserExt};

pub struct ProcessRes {
    pid:        u32,
    ppid:       u32,
    status:     String,
    username:   String,
    path:       String,
    command:    String,
    cwd:        String,
    environ:    String,
    name:       String,
}

impl ToString for ProcessRes {
    #[inline]
    fn to_string(&self) -> String {
        return format!("{{pid:{},ppid:{},status:\"{}\",username:\"{}\",path:\"{}\",\
        command:\"{}\",cwd:\"{}\",environ:\"{}\",name:\"{}\"}}",
        &self.pid,
        &self.ppid,
        &self.status,
        &self.username,
        &self.path,
        &self.command,
        &self.cwd,
        &self.environ,
        &self.name,
    );
    }
}

pub fn list() -> Result<Vec<String>> {
    if !System::IS_SUPPORTED {
        return Err(anyhow::anyhow!("This OS isn't supported for process functions.
         Pleases see sysinfo docs for a full list of supported systems.
         https://docs.rs/sysinfo/0.23.5/sysinfo/index.html#supported-oses\n\n"));
    }

    let mut res : Vec<String> = Vec::new();
    let mut sys = System::new();
    sys.refresh_processes();
    sys.refresh_users_list();
    let user_list =  sys.users().clone();

    for (pid, process) in sys.processes() {
        let mut tmp_ppid = 0;
        if  process.parent() != None {
            tmp_ppid = process.parent().unwrap().as_u32();
        }

        let tmprow = ProcessRes{
            pid:        pid.as_u32(),
            ppid:       tmp_ppid,
            status:     process.status().to_string(),
            username:   uid_to_username(process.uid, user_list),
            path:       String::from(process.exe().to_str().unwrap()),
            command:    String::from(process.cmd().join(" ")),
            cwd:        String::from(process.cwd().to_str().unwrap()),
            environ:    String::from(process.environ().join(" ")),
            name:       String::from(process.name())
        };
        res.push(tmprow.to_string());
    }
    Ok(res)
}

fn uid_to_username(username: u32, user_list: &[User]) -> String {
    for user in user_list {
        if *user.uid() == username {
            return String::from(user.name());
        }
    }
    return String::from("?");
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::process::Command;

    #[test]
    fn test_process_list() -> anyhow::Result<()>{
        let child = Command::new("sleep")
            .arg("5")
            .spawn()?;
        
        println!("{:?}", child.id());
        let res = list()?;
        let searchstring = String::from(format!("pid:{}", child.id()));
        for proc in res{
            if proc.as_str().contains(&searchstring) {
                println!("{:?}", proc.to_string());
                if proc.as_str().contains("command:\\\"sleep 5\\\"") {
                    assert_eq!(true, true);
                }
                return Ok(())
            }
        }
        assert_eq!(true, false);
        return Ok(())
    }
}